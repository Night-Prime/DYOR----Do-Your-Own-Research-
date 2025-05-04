package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/go-chi/chi/v5"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/routes"
)

func main() {

	_, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	cfg := config.Get()
	db := config.LoadDB()

	fmt.Printf("Successfully connected to database: %v\n", db)
	fmt.Println("--------------------------------------------- \n")

	fmt.Println("Currently Initializing Server")
	fmt.Println("--------------------------------------------- \n")

    // Set up CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
        AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by browsers
    })

	// Initialize the router
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	router.Use(c.Handler)

	apiRouter.Mount("/asset", routes.AssetRouteHandler())
	apiRouter.Mount("/admin", routes.AdminRouteHandler())
	apiRouter.Mount("/user", routes.UserRouteHandler())
	router.Mount("/api/v1", apiRouter)


	app := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("---------------------------------------------")
    fmt.Printf(" Starting DYOR Server on port: %s\n", cfg.Port)
	fmt.Println("---------------------------------------------")

	if err := app.ListenAndServe(); err != nil {
		fmt.Printf("Server failed to start: %v", err)
   	}
}