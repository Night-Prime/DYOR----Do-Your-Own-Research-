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

func (s *AssetService) GetAsset(assetType models.AssetType, symbols ...string) (*models.Asset, error) {
	fmt.Println("The Asset Service Layer")
	fmt.Println("--------------------------------------------- \n")

	var asset *models.Asset
	var err error

	switch assetType {
	case models.AssetTypeCrypto:
		if len(symbols) == 0 {
            return nil, fmt.Errorf("symbols are required for fetching crypto data")
        }
        asset, err = s.fetchCrypto(symbols)
	case models.AssetTypeStock:
		if len(symbols) == 0 {
			return nil, fmt.Errorf("symbol is required for fetching stock data")
		}
		symbol := symbols[0]
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
func (s *AssetService) fetchCrypto(symbols []string) (*models.Asset, error) {
	fmt.Println("Getting the data from the Crypto API Client Layer")
	fmt.Println("--------------------------------------------- \n")

	cryptoData, err := s.cryptoAPIClient.GetCryptoData(symbols)
	if err != nil {
		return nil, fmt.Errorf("error fetching crypto data: %v", err)
	}

	if len(cryptoData.DataArray) == 0 {
		return nil, fmt.Errorf("no crypto data found")
	}
	
	// Map through every index and create a list of assets
	var assets []*models.Asset
	for _, result := range cryptoData.DataArray {
		asset := &models.Asset{
			AssetBase: models.AssetBase{
				Type:         models.AssetTypeCrypto,
				Symbol:       result.Symbol,
				Name:         result.Name,
				CurrentPrice: result.Price, 
				Volume:       result.Volume,
			},
			CryptoData: cryptoData,
		}
		assets = append(assets, asset)
	}

	// Return the first asset for now or modify the function to return all assets
	if len(assets) == 0 {
		return nil, fmt.Errorf("no crypto data found")
	}
	return assets[0], nil
}