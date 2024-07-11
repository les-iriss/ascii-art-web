package main

import (
	"fmt"
	"log"
	"net/http"

	controller "ascii-art-web/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.GetRequest)
	mux.HandleFunc("/ascii-art", controller.PostRequest)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	PORT := ":8000"
	fmt.Printf("Server running on http://localhost%s\n", PORT)

	if err := http.ListenAndServe(PORT, mux); err != nil {
		log.Fatalf("Failed to start server on %s: %v", PORT, err)
	}
}
