package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/service"
)

func UserRouteHandler() http.Handler {
	newsClient := service.NewsClient()
	newsService := service.NewNewsService(newsClient)
	newsHandlers := handlers.NewNewsHandler(newsService)

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.UserAuthMiddleware)
		r.Get("/portfolio", handlers.GetPortfolioForUserHandler)
		r.Post("/portfolio", handlers.CreatePortfolioHandler)
		r.Put("/portfolio", handlers.UpdatePortfolioHandler)
		r.Delete("/portfolio", handlers.DeletePortfolioHandler)
		r.Get("/news", newsHandlers.GetNewsHandler)
		r.Get("/sentiment-score", newsHandlers.CalculateNewsSentimentHandler)
		r.Get("/top-gainers-losers", newsHandlers.GetTopGainersLosersHandler)
		r.Get("/assets-file", handlers.GetJSONFileHandler)
	})

	router.Post("/signup", handlers.SignupHandler)
	router.Post("/login", handlers.LoginHandler)
	router.Post("/logout", handlers.LogoutHandler)
	router.Get("/verify", handlers.VerifyUser)
	return router
}
