package main

import (
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.ConnStr))
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	fmt.Printf("Successfully connected to database: %v\n", db)
	fmt.Println("--------------------------------------------- \n")

	fmt.Println("Currently Initializing Server")
	fmt.Println("--------------------------------------------- \n")


	app := &http.Server{
		Addr: 			":" + cfg.Port,
		Handler: 		chi.NewRouter(),
		ReadTimeout:	60 * time.Second,
		WriteTimeout:	60 * time.Second,
	}

	fmt.Println("---------------------------------------------")
    fmt.Printf(" Starting DYOR Server on port: %s\n", cfg.Port)
	fmt.Println("---------------------------------------------")

	if err := app.ListenAndServe(); err != nil {
		fmt.Printf("Server failed to start: %v", err)
   	}
}