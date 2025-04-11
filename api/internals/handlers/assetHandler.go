package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)


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
	symbol := r.URL.Query().Get("symbols")

	// parse the asset type
	assetType := models.AssetType(assetTypeStr)

	// Call the asset service to get the asset data
	asset, err := h.assetService.GetAsset(assetType, symbol)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(asset)
}