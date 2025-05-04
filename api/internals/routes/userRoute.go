package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/middleware"
)

func UserRouteHandler() http.Handler {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.UserAuthMiddleware)
		r.Get("/portfolio", handlers.GetPortfolioForUserHandler)
		r.Post("/portfolio", handlers.CreatePortfolioHandler)
		r.Delete("/portfolio", handlers.DeletePortfolioHandler)
	})

	router.Post("/signup", handlers.SignupHandler)
	router.Post("/login", handlers.LoginHandler)
	router.Get("/verify", handlers.VerifyUser)
	return router
}
