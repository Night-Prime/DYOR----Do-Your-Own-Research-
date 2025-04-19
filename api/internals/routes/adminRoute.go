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
		r.Use(middleware.AuthMiddleware)
		r.Get("/", handlers.GetUserByEmailHandler)
		r.Get("/id", handlers.GetUserByIDHandler)
		r.Get("/all", handlers.GetAllUsersHandler)
		r.Put("/update", handlers.UpdateUserHandler)
		r.Delete("/delete", handlers.DeleteUserHandler)
	})

	router.Post("/signup", handlers.SignupHandler)
	router.Post("/login", handlers.LoginHandler)
	return router
}
