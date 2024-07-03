package main

import (
	"fmt"
	"log"
	"net/http"

	controller "ascii-art-web/controllers"
)

func main() {
	http.HandleFunc("/", controller.GetRequest)
	http.HandleFunc("/ascii-art", controller.PostRequest)

	PORT := ":8000"
	fmt.Printf("Server running on http://localhost%s\n", PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatalf("Failed to start server on %s: %v", PORT, err)
	}
}
