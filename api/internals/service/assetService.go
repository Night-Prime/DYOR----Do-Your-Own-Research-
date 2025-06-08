package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
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

func (s *AssetService) GetAsset(assetType models.AssetType, symbols ...string) (*models.Asset, error) {
	fmt.Println("The Asset Service Layer")
	fmt.Println("--------------------------------------------- \n")

	var asset *models.Asset
	var err error

	switch assetType {
	case models.AssetTypeCrypto:
		if len(symbols) == 0 {
            return nil, fmt.Errorf("Symbols are required for fetching crypto data")
        }
        asset, err = s.fetchCrypto(symbols)
	case models.AssetTypeStock:
		if len(symbols) == 0 {
			return nil, fmt.Errorf("Symbol is required for fetching stock data")
		}
		symbol := symbols[0]
		asset, err = s.fetchStock(symbol)
	// case models.AssetTypeBond:
	// 	asset, err = s.fetchBond()
	default:
		return nil, fmt.Errorf("Unsupported asset type: %s", assetType)
	}

	if err != nil {
		return nil, fmt.Errorf("Error fetching asset data: %v", err)
	}

	if err := asset.Validate(); err != nil {
		return nil, fmt.Errorf("Asset validation error: %v", err)
	}

	return asset, nil
}

// For Stocks:
func (s *AssetService) fetchStock(symbol string) (*models.Asset, error) {
    fmt.Println("Getting the data from the Stock API Client Layer", symbol)
    fmt.Println("--------------------------------------------- \n")

    stockData, err := s.stockAPIClient.GetStockData(symbol)
    if err != nil {
        return nil, fmt.Errorf("Error fetching stock data: %v", err)
    }

    // Check if we have results
    if len(stockData.Data.QuoteResponse.Result) == 0 {
        return nil, fmt.Errorf("No stock data found for symbol: %s", symbol)
    }
    result := stockData.Data.QuoteResponse.Result[0]

    // Create the Asset with properly mapped fields
    asset := &models.Asset{
        AssetBase: models.AssetBase{
            Type:         models.AssetTypeStock,
            Symbol:       symbol,
            CurrentPrice: result.RegularMarketPrice.Raw,
            Volume:       result.RegularMarketVolume.Raw,
            Name:         result.Name,
        },
        StockData: stockData,
    }
    
    return asset, nil
}

// For Crypto :
func (s *AssetService) fetchCrypto(symbols []string) (*models.Asset, error) {
	fmt.Println("Getting the data from the Crypto API Client Layer")
	fmt.Println("--------------------------------------------- \n")

	cryptoData, err := s.cryptoAPIClient.GetCryptoData(symbols)
	if err != nil {
		return nil, fmt.Errorf("Error fetching crypto data: %v", err)
	}

	if len(cryptoData.DataArray) == 0 {
		return nil, fmt.Errorf("No crypto data found")
	}
	
    asset := &models.Asset{
        AssetBase: models.AssetBase{
            Type: models.AssetTypeCrypto,
        },
        CryptoData: cryptoData,
    }

	return asset, nil
}