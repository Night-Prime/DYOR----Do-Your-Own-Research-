package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

// (Another DI happening here):
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