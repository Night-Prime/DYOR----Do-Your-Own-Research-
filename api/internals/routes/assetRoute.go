package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

func AssetRouteHandler() http.Handler {
	stockClient := service.NewStockClient()
	cryptoClient := service.NewCryptoClient()
	// bondClient := service.NewBondClient()
	
	// Create service with all dependencies
	assetService := service.NewAssetService(stockClient, cryptoClient)
	
	// Create handler
	assetHandler := handlers.NewAssetHandler(assetService)

	router := chi.NewRouter()
	router.Get("/", assetHandler.GetAssetHandler)
	
	return router
}