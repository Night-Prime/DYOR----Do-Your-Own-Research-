package service

import (
	"fmt"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
)

// The Unified Asset Service: (A lot of dependency injection happening here too)
// This service is responsible for processing asset.
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

func (s *AssetService) GetAsset(assetType models.AssetType, args ...string) (*models.Asset, error) {
	fmt.Println("The Asset Service Layer")
	fmt.Println("--------------------------------------------- \n")

	var asset *models.Asset
	var err error

	switch assetType {
	case models.AssetTypeCrypto:
		var page, currency, per_page string

		if len(args) > 0 {
			if len(args) >= 1 {
				page = args[0]
			}
			if len(args) >= 2 {
				currency = args[1]
			}
			if len(args) >= 3 {
				per_page = args[2]
			}
		}

		asset, err = s.fetchCrypto(page, currency, per_page)
	case models.AssetTypeStock:
		if len(args) == 0 {
			return nil, fmt.Errorf("symbol is required for fetching stock data")
		}
		symbol := args[0]
		asset, err = s.fetchStock(symbol)
	// case models.AssetTypeBond:
	// 	asset, err = s.fetchBond()
	default:
		return nil, fmt.Errorf("unsupported asset type: %s", assetType)
	}

	if err != nil {
		return nil, fmt.Errorf("error fetching asset data: %v", err)
	}

	if err := asset.Validate(); err != nil {
		return nil, fmt.Errorf("asset validation error: %v", err)
	}

	return asset, nil
}

// For Stocks:
func (s *AssetService) fetchStock(symbol string) (*models.Asset, error) {
    fmt.Println("Getting the data from the Stock API Client Layer", symbol)
    fmt.Println("--------------------------------------------- \n")

    stockData, err := s.stockAPIClient.GetStockData(symbol)
    if err != nil {
        return nil, fmt.Errorf("error fetching stock data: %v", err)
    }

    // Check if we have results
    if len(stockData.Data.QuoteResponse.Result) == 0 {
        return nil, fmt.Errorf("no stock data found for symbol: %s", symbol)
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
func (s *AssetService) fetchCrypto(page, currency, per_page string) (*models.Asset, error) {
	fmt.Println("Getting the data from the Crypto API Client Layer")
	fmt.Println("--------------------------------------------- \n")

	cryptoData, err := s.cryptoAPIClient.GetCryptoData(page, currency, per_page)
	if err != nil {
		return nil, fmt.Errorf("error fetching crypto data: %v", err)
	}

	if len(*cryptoData) == 0 {
		return nil, fmt.Errorf("no crypto data found")
	}

	// Iterate over the array and process each object
	for _, result := range *cryptoData {
		asset := &models.Asset{
			AssetBase: models.AssetBase{
				Type:         models.AssetTypeCrypto,
				Symbol:       result.Symbol, // Replace with the correct field
				Volume:       result.CirculatingSupply, // Replace with the correct field
				CurrentPrice: result.High24H, // Replace with the correct field
				Name:         result.LastUpdated, // Replace with the correct field
			},
			CryptoData: cryptoData,
		}
		
		return asset, nil
	}

	return nil, fmt.Errorf("no valid crypto asset found")
}