package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
)

func AssetRouteHandler() http.Handler {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.UserAuthMiddleware)
		r.Post("/create-asset", handlers.CreateAssetHandler)
		r.Delete("/delete-asset", handlers.DeleteAssetHandler)
	})

	return router
}

// func AssetRouteHandler() http.Handler {
// 	stockClient := service.NewStockClient()
// 	cryptoClient := service.NewCryptoClient()
// 	// bondClient := service.NewBondClient()
	
// 	assetService := service.NewAssetService(stockClient, cryptoClient)
	
// 	// Create handler
// 	assetHandler := handlers.NewAssetHandler(assetService)

// 	router := chi.NewRouter()
// 	router.Get("/get-live-update", assetHandler.GetAssetHandler)
	
// 	return router
// }

