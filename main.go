package main

import (
	routes "css/mlc/routes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// os.Setenv("SPOTIFY_ID", "e3e3b5ede9344cebaecd3d44b669ea6f")
	// os.Setenv("SPOTIFY_SECRET", "0e312d906a554309b28a7297162f1bfc")
	// fmt.Println(os.LookupEnv("SPOTIFY_SECRET"))
	// fmt.Println(os.LookupEnv("SPOTIFY_ID"))
	port := "127.0.0.1:1234"
	runServer(port)
}

func runServer(port string) {
	r := routes.Start_Routes()
	fmt.Printf("Listening on port %s\n", port)
	server := http.Server{
		Addr:    port,
		Handler: r,
	}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
