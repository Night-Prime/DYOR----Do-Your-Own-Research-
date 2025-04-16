package routes

import (
	"net/http"
	"github.com/go-chi/chi/v5"

	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/handlers"
)

func UserRouteHandler() http.Handler {
	router := chi.NewRouter()
	router.Get("/", handlers.GetUserByEmailHandler)
	router.Get("/id", handlers.GetUserByIDHandler)
	router.Get("/all", handlers.GetAllUsersHandler)
	router.Post("/signup", handlers.CreateUserHandler)
	router.Put("/update", handlers.UpdateUserHandler)
	router.Delete("/delete", handlers.DeleteUserHandler)
	
	return router
}
