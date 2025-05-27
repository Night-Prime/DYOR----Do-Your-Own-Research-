package main

import (
	"fmt"
	"net/http"
	"time"
	"os/signal"
	"os"
	"syscall"
	"log"
	"context"

	// "net/http/pprof"
	"github.com/rs/cors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/routes"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/database"
)

func main() {

	_, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	cfg := config.Get()

	if err := database.HealthCheck(); err != nil {
		log.Fatalf("Database health check failed: %v", err)
	}

	fmt.Printf("Successfully connected to database \n")
	fmt.Println("--------------------------------------------- \n")

	fmt.Println("Currently Initializing Server")
	fmt.Println("--------------------------------------------- \n")


	// Initialize the router
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	router.Use(
		middleware.Recoverer,
		cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Cookie"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}).Handler,
		middleware.Heartbeat("/health"),
	)

	apiRouter.Mount("/asset", routes.AssetRouteHandler())
	apiRouter.Mount("/admin", routes.AdminRouteHandler())
	apiRouter.Mount("/user", routes.UserRouteHandler())
	router.Mount("/api/v1", apiRouter)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// Setting up Channel Listeners:
	serverErrors := make(chan error, 1)
    shutdown := make(chan os.Signal, 1)
    
    signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

    go func() {
        fmt.Println("---------------------------------------------")
        fmt.Printf(" Starting DYOR Server on port: %s\n", cfg.Port)
        fmt.Println("---------------------------------------------")
        serverErrors <- server.ListenAndServe()
    }()

    select {
    case err := <-serverErrors:
        log.Fatalf("Server failed to start: %v", err)

    case sig := <-shutdown:
        log.Printf("Shutdown signal received: %v", sig)
        
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel() // Important to call cancel to release resources

        // Attempt graceful shutdown
        if err := server.Shutdown(ctx); err != nil {
            log.Printf("Graceful shutdown failed: %v", err)
            if err := server.Close(); err != nil {
                log.Fatalf("Force shutdown failed: %v", err)
            }
        }
    }

    log.Println("Server stopped gracefully")
}