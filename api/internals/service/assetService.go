package service

import (
	"fmt"
    "encoding/json"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/utils"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
)

// Note: didn't use DI on parts of the code not interacting with external services

func CreateAsset(assetType models.AssetType, symbolMap map[string]string, portfolioID uuid.UUID) ([]*models.Asset, error) {
    fmt.Println("Creating assets in the Asset Service Layer")
    fmt.Println("---------------------------------------------\n")

    var assets []*models.Asset

    for symbol, name := range symbolMap {
        var asset *models.Asset
        
        if name == "" {
            name = symbol
        }

        switch assetType {
        case models.AssetTypeStock:
            asset = &models.Asset{
                AssetBase: models.AssetBase{
                    Type:         models.AssetTypeStock,
                    Symbol:       symbol,
                    Name:         name,
                    PortfolioID:  portfolioID,
                    Quantity:     0,
                    CurrentPrice: 0,
                    Volume:       0,
                },
                StockData: nil,
            }
        case models.AssetTypeCrypto:
            asset = &models.Asset{
                AssetBase: models.AssetBase{
                    Type:         models.AssetTypeCrypto,
                    Symbol:       symbol,
                    Name:         name,
                    PortfolioID:  portfolioID,
                    Quantity:     0,
                    CurrentPrice: 0,
                    Volume:       0,
                },
                CryptoData: nil,
            }
        default:
            return nil, fmt.Errorf("Unsupported asset type: %s", assetType)
        }
        assets = append(assets, asset)
    }

    if err := models.SaveAssetToDB(assets); err != nil {
        return nil, err
    }

    return assets, nil
}

func DeleteAsset(assetID string) error {
	if err := models.DeleteAsset(assetID); err != nil {
		return err
	}
	return nil
}

// The Unified Asset Service section: (A lot of dependency injection happening here too)
// This part of the service is responsible for processing asset.
// It abstracts the complexity of dealing with multiple APIs and provides a unified interface for asset data retrieval.
// It handles different asset types such as stocks, bonds, and cryptocurrencies.

type AssetService struct {
	stockAPIClient StockAPIClient
	cryptoAPIClient CryptoAPIClient
}

func NewAssetService(stockAPIClient StockAPIClient, cryptoAPIClient CryptoAPIClient) *AssetService {
	return &AssetService{
		stockAPIClient: stockAPIClient,
		cryptoAPIClient: cryptoAPIClient,
	}
}

func (s *AssetService) GetAssets(assetType models.AssetType, portfolio *models.PortfolioReq, symbols ...string) ([]*models.Asset, error) {
    cfg := config.Get()
    switch assetType {
    case models.AssetTypeStock:
        return s.fetchStocks(symbols, cfg.REALTIME, portfolio)
    case models.AssetTypeCrypto:
        return s.fetchCrypto(symbols, cfg.REALTIME, portfolio)
    default:
        return nil, fmt.Errorf("Unsupported asset type: %s", assetType)
    }
}

func (s *AssetService) fetchStocks(symbols []string, allowRealTime bool, portfolio *models.PortfolioReq) ([]*models.Asset, error) {
    var assets []*models.Asset
    // TODO : Introduce concurrency approach to handle multi-request
    if allowRealTime {
        for _, symbol := range symbols {
            stockDataList, err := s.stockAPIClient.GetStockData(symbol)
            if err != nil {
                return nil, err
            }
            
            for _, stockData := range stockDataList {
                assets = append(assets, &models.Asset{
                    AssetBase: models.AssetBase{
                        Symbol: stockData.Symbol,
                        Name:   stockData.Name,
                        PortfolioID:  portfolio.PortfolioID,
                        Type:   models.AssetTypeStock,
                    },
                    StockData: &stockData,
                })

                if err := models.SaveAssetToDB(assets); err != nil {
                    return nil, err
                }
            }
        }
        
        if len(assets) == 0 {
            return nil, &errors.ValidationError{
                Message: "No stock found",
            }
        }
        
    } else {
        // portfolio, err := &models.GetPortfolioAssets(portfolio.PortfolioID)
        // if err != nil{
        //     return nil, err
        // }

        // assets =
    }

    return assets, nil
}

func (s *AssetService) fetchCrypto(symbols []string, allowRealTime bool, portfolio *models.PortfolioReq) ([]*models.Asset, error) {
    var assets []*models.Asset
    
    if allowRealTime {
        cryptoDataList, err := s.cryptoAPIClient.GetCryptoData(symbols)
        if err != nil {
            return nil, err
        }
        
        for _, cryptoData := range cryptoDataList {
            assets = append(assets, &models.Asset{
                AssetBase: models.AssetBase{
                    Symbol: cryptoData.Symbol,
                    Name:   cryptoData.Name,
                    PortfolioID:  portfolio.PortfolioID,
                    Type:   models.AssetTypeCrypto,
                },
                CryptoData: &cryptoData,
            })

            if err := models.SaveAssetToDB(assets); err != nil {
                return nil, err
            }
        }
        
        if len(assets) == 0 {
            return nil, &errors.ValidationError{
                Message: "No crypto found",
            }
        }
        
    } else {
        return assets, nil
    }

    return assets, nil
}

// AI-related Service : 

type AIService struct {
    aiInsightClient AIInsightClient
    stockAPIClient StockAPIClient
	cryptoAPIClient CryptoAPIClient
}

func NewAIService(aiInsightClient AIInsightClient, stockAPIClient StockAPIClient, cryptoAPIClient CryptoAPIClient) *AIService{
    return &AIService{
        aiInsightClient : aiInsightClient,
        stockAPIClient: stockAPIClient,
		cryptoAPIClient: cryptoAPIClient,
    }
}


func (a *AIService) GetAIInsightsSummary(asset string, promptType string, allowRealTime bool, assetType string) (models.FinancialAnalysisResponse, error) {
    var assetData []*models.Asset
    var formattedData string
    
    if allowRealTime {     
        fmt.Println("Is anything happening here ?")   
        switch assetType {
        case "crypto":
            arrAsset := []string{asset}
            cryptoDataList, err := a.cryptoAPIClient.GetCryptoData(arrAsset)
            if err != nil {
                return models.FinancialAnalysisResponse{}, fmt.Errorf("Failed to get crypto data: %w", err)
            }

            assetData = make([]*models.Asset, 0, len(cryptoDataList))
            for _, cryptoData := range cryptoDataList {
                assetData = append(assetData, &models.Asset{
                    AssetBase: models.AssetBase{
                        Symbol: cryptoData.Symbol, // Use actual symbol from data
                        Name:   cryptoData.Name,
                        Type:   models.AssetTypeCrypto,
                    },
                    CryptoData: &cryptoData,
                })
            }

        case "stock":
            stockDataList, err := a.stockAPIClient.GetStockData(asset)
            if err != nil {
                return models.FinancialAnalysisResponse{}, fmt.Errorf("Failed to get stock data: %w", err)
            }

            assetData = make([]*models.Asset, 0, len(stockDataList))
            for _, stockData := range stockDataList {
                assetData = append(assetData, &models.Asset{
                    AssetBase: models.AssetBase{
                        Symbol: stockData.Symbol, // Use actual symbol from data
                        Name:   stockData.Name,
                        Type:   models.AssetTypeStock,
                    },
                    StockData: &stockData,
                })
            }

        default:
            return models.FinancialAnalysisResponse{}, fmt.Errorf("Unsupported asset type: %s", assetType)
        }

        if len(assetData) == 0 {
            return models.FinancialAnalysisResponse{}, fmt.Errorf("No asset data available")
        }
    
        formattedLiveData, err := json.Marshal(assetData)
        if err != nil {
            return models.FinancialAnalysisResponse{}, fmt.Errorf("Failed to marshal asset data: %w", err)
        }

        formattedData = string(formattedLiveData)
        fmt.Printf("The Formatted Data: ", formattedData);
    } else {
        formattedData = asset
    }

    analysis, err := a.aiInsightClient.GetAssetInsight(formattedData, promptType)
    if err != nil {
        return models.FinancialAnalysisResponse{}, fmt.Errorf("Failed to get AI insights: %w", err)
    }

    plainText := utils.StripAllFormatting(analysis.Sentiment)
    analysis.Sentiment = plainText
    fmt.Printf("Here: %v", analysis)
    return analysis, nil
}