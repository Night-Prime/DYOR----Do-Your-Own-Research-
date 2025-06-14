package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

// Note: didn't use DI on parts of the code not interacting with external services

func CreateAsset(assetType models.AssetType, symbolMap map[string]string, portfolioID uuid.UUID) ([]*models.Asset, error) {
    fmt.Println("Creating multiple assets in the Asset Service Layer")
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
	// bondAPIClient BondAPIClient
}

func NewAssetService(stockAPIClient StockAPIClient, cryptoAPIClient CryptoAPIClient) *AssetService {
	return &AssetService{
		stockAPIClient: stockAPIClient,
		cryptoAPIClient: cryptoAPIClient,
		// bondAPIClient: bondAPIClient,
	}
}

func (s *AssetService) GetAssets(assetType models.AssetType, symbols ...string) ([]*models.Asset, error) {
    switch assetType {
    case models.AssetTypeStock:
        return s.fetchStocks(symbols)
    case models.AssetTypeCrypto:
        return s.fetchCrypto(symbols)
    default:
        return nil, fmt.Errorf("unsupported asset type: %s", assetType)
    }
}

func (s *AssetService) fetchStocks(symbols []string) ([]*models.Asset, error) {
    var assets []*models.Asset
    
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
                    Type:   models.AssetTypeStock,
                },
                StockData: &stockData,
            })
        }
    }
    
    if len(assets) == 0 {
        return nil, &errors.ValidationError{
            Message: "No stock found",
        }
    }
    
    return assets, nil
}

func (s *AssetService) fetchCrypto(symbols []string) ([]*models.Asset, error) {
    cryptoDataList, err := s.cryptoAPIClient.GetCryptoData(symbols)
    if err != nil {
        return nil, err
    }
    
    var assets []*models.Asset
    for _, cryptoData := range cryptoDataList {
        assets = append(assets, &models.Asset{
            AssetBase: models.AssetBase{
                Symbol: cryptoData.Symbol,
                Name:   cryptoData.Name,
                Type:   models.AssetTypeCrypto,
            },
            CryptoData: &cryptoData,
        })
    }
    
    if len(assets) == 0 {
        return nil, &errors.ValidationError{
            Message: "No crypto found",
        }
    }
    
    return assets, nil
}