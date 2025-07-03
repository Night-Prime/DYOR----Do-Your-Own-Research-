package service

import (
	"fmt"
    "strings"
	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
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
    // TODO : Introduce concurrency approach to handle multi-request
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

// AI-related Service : 

type AIService struct {
    aiInsightClient AIInsightClient
}

func NewAIService(aiInsightClient AIInsightClient) *AIService{
    return &AIService{
        aiInsightClient : aiInsightClient,
    }
}


func (a *AIService) GetAIInsightsSummary(asset string) (string, error) {
    prompt := fmt.Sprintf(`{
        "test": "So I'm investing into $%s, I need a breakdown of the financial statements & performance over the past year"
    }`, asset)
    
    analysis, err := a.aiInsightClient.GetAssetInsight(prompt)
    if err != nil {
        return "", &errors.DatabaseError{
            Message: "Error Fetching Data from the AI Service",
            Err: err,
        }
    }

    // Helper functions
    getFirstPriceTarget := func(targets map[string]float64) float64 {
        for _, v := range targets {
            return v
        }
        return 0
    }

    getFirstProbability := func(probs map[string]string) string {
        for _, v := range probs {
            return v
        }
        return "N/A"
    }


    // Get first support/resistance level if available
    getFirstSupport := func() string {
        if len(analysis.FinancialAnalysis.TechnicalAnalysis.SupportLevels) > 0 {
            return fmt.Sprintf("%.0f", analysis.FinancialAnalysis.TechnicalAnalysis.SupportLevels[0])
        }
        return "not specified"
    }

    getFirstResistance := func() string {
        if len(analysis.FinancialAnalysis.TechnicalAnalysis.ResistanceLevels) > 0 {
            return fmt.Sprintf("%.0f", analysis.FinancialAnalysis.TechnicalAnalysis.ResistanceLevels[0])
        }
        return "not specified"
    }

    formattedSummary := fmt.Sprintf(`
    Based on the latest financial analysis, the overall market sentiment is currently %s with a confidence level of %s, primarily driven by factors such as %s. 
    From a technical standpoint, the market is showing signs of a %s, with key support and resistance levels identified for major assets. For instance, the asset has support around %s and resistance at %s. 
    Key indicators suggest continued strength in several assets, particularly in the technology sector. 
    
    Risk assessment indicates a %s risk level, with major concerns including %s. Volatility is assessed as *%s*, suggesting investors should remain cautiously optimistic.
    
    As for actionable insights, the analysis recommends: %s. The projected price target is $%.0f within a time horizon of %s. The likelihood of hitting this target is %s.
    
    This analysis was generated by model %s on %s, offering a balanced blend of market signals, technical evaluation, and actionable strategies tailored for medium-term investors.
    `,
        analysis.FinancialAnalysis.MarketSentiment.OverallSentiment,
        analysis.FinancialAnalysis.MarketSentiment.ConfidenceLevel,
        strings.Join(analysis.FinancialAnalysis.MarketSentiment.KeyDrivers, ", "),
        analysis.FinancialAnalysis.TechnicalAnalysis.Trend,
        getFirstSupport(),
        getFirstResistance(),
        analysis.FinancialAnalysis.RiskAssessment.RiskLevel,
        strings.Join(analysis.FinancialAnalysis.RiskAssessment.KeyRisks, ", "),
        analysis.FinancialAnalysis.RiskAssessment.VolatilityAssessment,
        strings.Join(analysis.FinancialAnalysis.ActionableInsights.Recommendations, "; "),
        getFirstPriceTarget(analysis.FinancialAnalysis.ActionableInsights.PriceTargets),
        analysis.FinancialAnalysis.ActionableInsights.TimeHorizon,
        getFirstProbability(analysis.FinancialAnalysis.ActionableInsights.ProbabilityAssessments),
        analysis.Metadata.Model,
        analysis.Metadata.Timestamp,
    )

    return formattedSummary, nil
}