package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

func AssetRouteHandler() http.Handler {
	stockClient := service.NewStockClient()
	cryptoClient := service.NewCryptoClient()
	aiClient := service.NewAIClient()

	assetService := service.NewAssetService(stockClient, cryptoClient)
	aiService := service.NewAIService(aiClient)
	assetHandler := handlers.NewAssetHandler(assetService, aiService)
	

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.UserAuthMiddleware)
		r.Post("/create-asset", handlers.CreateAssetsHandler)
		r.Delete("/delete-asset", handlers.DeleteAssetHandler)
		r.Get("/get-live-update", assetHandler.GetAssetHandler)
		r.Get("/get-ai-summary", assetHandler.GetAIInsightsHandler)
	})

	return router
}
