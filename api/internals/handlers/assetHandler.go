package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/models"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
	"github.com/google/uuid"
)

func CreateAssetHandler(w http.ResponseWriter, r *http.Request) {
	assetRequest := &models.Asset{}
	if err := json.NewDecoder(r.Body).Decode(&assetRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	portfolioID, err := uuid.Parse(assetRequest.PortfolioID.String())
	if err != nil {
		http.Error(w, "Invalid portfolio ID format", http.StatusBadRequest)
		return
	}

	assetType := models.AssetType(assetRequest.Type)
	if assetType != models.AssetTypeStock && assetType != models.AssetTypeCrypto {
		http.Error(w, "Invalid asset type", http.StatusBadRequest)
		return
	}

	asset, err := service.CreateAsset(assetType, assetRequest.Symbol, portfolioID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asset)
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

type AssetHandler struct {
	assetService *service.AssetService
}

func NewAssetHandler(assetService *service.AssetService) *AssetHandler {
	return &AssetHandler{
		assetService: assetService,
	}
}

func (h *AssetHandler) GetAssetHandler(w http.ResponseWriter, r *http.Request) {
	assetTypeStr := r.URL.Query().Get("type")
	assetType := models.AssetType(assetTypeStr)

	var asset interface{}
	var err error

	switch assetType {
	case "stock":
		symbol := r.URL.Query().Get("symbols")
		if symbol == "" {
			http.Error(w, "symbol is required for stock type", http.StatusBadRequest)
			return
		}
		asset, err = h.assetService.GetAsset(assetType, symbol)
	case "crypto":
		symbols := r.URL.Query()["symbols"]
		if len(symbols) == 0 {
			http.Error(w, "at least one symbol is required for crypto type", http.StatusBadRequest)
			return
		}

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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asset)
}