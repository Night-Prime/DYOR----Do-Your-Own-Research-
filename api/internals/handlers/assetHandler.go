package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
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
        http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ( DI happening here, the intention here is to connect to an external service/ model to get insights on each assets):
type AssetHandler struct {
	assetService  *service.AssetService
}

func NewAssetHandler(assetService *service.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService: assetService,
	}
}


func (h *AssetHandler) GetAssetHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the asset type from the request
	assetTypeStr := r.URL.Query().Get("type")

	// parse the asset type
	assetType := models.AssetType(assetTypeStr)

	var asset interface{}
	var err error

	switch assetType {
	case "stock":
		symbol := r.URL.Query().Get("symbols")
		asset, err = h.assetService.GetAsset(assetType, symbol)
	case "crypto":
		// Get all symbol values from query params
		symbols := r.URL.Query()["symbols"]
		if len(symbols) == 0 {
			http.Error(w, "at least one symbol is required", http.StatusBadRequest)
			return
		}
		
		// Convert to variadic arguments
		args := make([]string, 0, len(symbols))
		args = append(args, symbols...)
		
		asset, err = h.assetService.GetAsset(assetType, args...)
	default:
		http.Error(w, "Invalid asset type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asset)
}