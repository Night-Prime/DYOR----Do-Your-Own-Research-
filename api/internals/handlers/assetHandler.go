package handlers

import (
    "strings"
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
)

// Note: didn't use DI on parts of the code not interacting with external services

func CreateAssetsHandler(w http.ResponseWriter, r *http.Request) {
    // Define request struct to match your payload
	type AssetSymbol struct {
        Symbol string `json:"symbol"`
        Name   string `json:"name"`
    }

    type AssetRequest struct {
        Type        string        `json:"type"`
        Symbols     []AssetSymbol `json:"symbols"`
        PortfolioID uuid.UUID     `json:"portfolioID"`
    }

    var request AssetRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	if len(request.Symbols) == 0 {
        http.Error(w, "At least one symbol is required", http.StatusBadRequest)
        return
    }

	// Convert to symbol map (symbol -> name)
    symbolMap := make(map[string]string)
    for _, symbol := range request.Symbols {
        if symbol.Symbol == "" {
            http.Error(w, "Symbol cannot be empty", http.StatusBadRequest)
            return
        }
        symbolMap[symbol.Symbol] = symbol.Name
    }

    // Validate asset type (come back to this)
    assetType := models.AssetType(request.Type)
    if assetType != models.AssetTypeStock && assetType != models.AssetTypeCrypto {
        http.Error(w, "Invalid asset type", http.StatusBadRequest)
        return
    }

    // Create the assets
    assets, err := service.CreateAsset(assetType, symbolMap, request.PortfolioID)
    if err != nil {
        switch err.(type) {
        case *errors.ValidationError:
            http.Error(w, err.Error(), http.StatusBadRequest) 
        case *errors.DatabaseError:
            http.Error(w, err.Error(), http.StatusInternalServerError) 
		case *errors.CustomError:
            http.Error(w, err.Error(), http.StatusInternalServerError)
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError) 
        }
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(assets)
}

func DeleteAssetHandler(w http.ResponseWriter, r *http.Request) {
	assetID := r.URL.Query().Get("id")
	if assetID == "" {
		http.Error(w, "Asset ID is required for deletion", http.StatusBadRequest)
		return
	}

	err := service.DeleteAsset(assetID)
	if err != nil {
        switch err.(type) {
        case *errors.ValidationError:
            http.Error(w, err.Error(), http.StatusBadRequest)
        case *errors.DatabaseError:
            http.Error(w, err.Error(), http.StatusInternalServerError)
		case *errors.CustomError:
            http.Error(w, err.Error(), http.StatusInternalServerError)
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ( DI happening here (i don't think it should get to this level), the intention here is to connect to an external service/ model to get insights on each assets):
type AssetHandler struct {
	assetService  *service.AssetService
}

func NewAssetHandler(assetService *service.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService: assetService,
	}
}

func (h *AssetHandler) GetAssetHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    
    response := struct {
        Stocks []*models.Asset `json:"stocks,omitempty"`
        Crypto []*models.Asset `json:"crypto,omitempty"`
        Errors []string        `json:"errors,omitempty"`
    }{}

    if stockSymbols := query.Get("stock_symbols"); stockSymbols != "" {
        symbols := strings.Split(stockSymbols, ",")
        stocks, err := h.assetService.GetAssets(models.AssetTypeStock, symbols...)
        if err != nil {
            switch err.(type) {
            case *errors.ValidationError:
                http.Error(w, err.Error(), http.StatusBadRequest) 
            case *errors.DatabaseError:
                http.Error(w, err.Error(), http.StatusInternalServerError) 
            default:
                http.Error(w, err.Error(), http.StatusInternalServerError) 
            }
            return
        }
        response.Stocks = stocks
    }
    
    if cryptoSymbols := query.Get("crypto_symbols"); cryptoSymbols != "" {
        symbols := strings.Split(cryptoSymbols, ",")
        cryptos, err := h.assetService.GetAssets(models.AssetTypeCrypto, symbols...)
        if err != nil {
            switch err.(type) {
            case *errors.ValidationError:
                http.Error(w, err.Error(), http.StatusBadRequest) 
            case *errors.DatabaseError:
                http.Error(w, err.Error(), http.StatusInternalServerError) 
            default:
                http.Error(w, err.Error(), http.StatusInternalServerError) 
            }
            return
        }
        response.Crypto = cryptos
    }

    
    if len(response.Stocks) == 0 && len(response.Crypto) == 0 && len(response.Errors) == 0 {
        http.Error(w, "No valid asset types or symbols provided", http.StatusBadRequest)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}