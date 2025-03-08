package main

import (
	"fmt"
	"os"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

func main() {
	godotenv.Load(".env")

	fmt.Println("Currrently Initializing Server")
	fmt.Println("--------------------------------------------- \n")

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT is not found")
	}
	port = ":" + port

	app := &http.Server{
		Addr: 			port,
		Handler: 		chi.NewRouter(),
		ReadTimeout:	60 * time.Second,
		WriteTimeout:	60 * time.Second,
	}

	fmt.Println("---------------------------------------------")
    fmt.Printf(" Starting DYOR Server on port%s\n", port)
	fmt.Println("---------------------------------------------")

	if err := app.ListenAndServe(); err != nil {
		fmt.Printf("Server failed to start: %v", err)
   	}
}