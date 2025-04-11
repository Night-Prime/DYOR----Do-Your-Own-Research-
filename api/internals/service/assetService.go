package service

import (
	"fmt"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
)

// The Unified Asset Service: (A lot of dependency injection happening in this codebase)
// This service is responsible for processing asset.
// It abstracts the complexity of dealing with multiple APIs and provides a unified interface for asset data retrieval.
// It handles different asset types such as stocks, bonds, and cryptocurrencies.

type AssetService struct {
	stockAPIClient StockAPIClient
	// cryptoAPIClient CryptoAPIClient
	// bondAPIClient BondAPIClient
}

func NewAssetService(stockAPIClient StockAPIClient) *AssetService {
	return &AssetService{
		stockAPIClient: stockAPIClient,
		// cryptoAPIClient: cryptoAPIClient,
		// bondAPIClient: bondAPIClient,
	}
}

func (s *AssetService) GetAsset(assetType models.AssetType, symbol string) (*models.Asset, error) {
	fmt.Println("The Asset Service Layer")
	fmt.Println("--------------------------------------------- \n")

	var asset *models.Asset
	var err error

	switch assetType {
	// case models.AssetTypeCrypto:
	// 	asset, err = s.fetchCrypto()
	case models.AssetTypeStock:
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

// func (s *AssetService) fetchCrypto() (*models.Asset, error) {
// 	cryptoData, err := s.cryptoAPIClient.GetCryptoData()
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching crypto data: %v", err)
// 	}

// 	// return &models.Asset{
// 	// 	AssetBase: models.AssetBase{
// 	// 		Type: models.AssetTypeCrypto,
// 	// 		Symbol: cryptoData.Symbol,
// 	// 		CurrentPrice: cryptoData.CurrentPrice,
// 	// 		Volume: cryptoData.Volume,
// 	// 		Quantity: cryptoData.Quantity,
// 	// 		Name: cryptoData.Name,
// 	// 	},
// 	// 	CryptoData: &models.CryptoData{
// 	// 		MarketCap: cryptoData.MarketCap,
// 	// 		CirculatingSupply: cryptoData.CirculatingSupply,
// 	// 		MaxSupply: cryptoData.MaxSupply,
// 	// 		LastUpdated: cryptoData.LastUpdated,
// 	// 	},
// 	// }, nil
// 	fmt.Printf("Model Data &v+", cryptoData)
// 	return &models.Asset{}, nil
// }

// func (s *AssetService) fetchBond() (*models.Asset, error) {
// 	bondData, err := s.bondAPIClient.GetBondData()
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching bond data: %v", err)
// 	}

// 	// return &models.Asset{
// 	// 	AssetBase: models.AssetBase{
// 	// 		Type: models.AssetTypeBond,
// 	// 		Symbol: bondData.Symbol,
// 	// 		CurrentPrice: bondData.CurrentPrice,
// 	// 		Volume: bondData.Volume,
// 	// 		Quantity: bondData.Quantity,
// 	// 		Name: bondData.Name,
// 	// 	},
// 	// 	BondData: &models.BondData{
// 	// 		YieldToMaturity: bondData.YieldToMaturity,
// 	// 		MaturityDate: bondData.MaturityDate,
// 	// 		CouponRate: bondData.CouponRate,
// 	// 		CreditRating: bondData.CreditRating,
// 	// 	},
// 	// }, nil
// 	fmt.Printf("Model Data &v+", bondData)
// 	return &models.Asset{}, nil
// }