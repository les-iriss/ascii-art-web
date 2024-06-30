package main

import (
	"fmt"
	"log"
	"net/http"

	controller "ascii-art-web/controllers"
)

func main() {
	

	http.HandleFunc("/", controller.GetPage)
	http.HandleFunc("/ascii-art", controller.PostRequest)
    fmt.Println("server running on 'http://localhost:8000'")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
	
}
