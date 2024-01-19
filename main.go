package main

import (
	routes "css/mlc/routes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	port := "127.0.0.1:3333"
	wg.Add(1)
	go runServer(port, &wg)
	wg.Wait()
}

func runServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	r := routes.Start_Routes()
	fmt.Printf("Listening on port %s", port)
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
