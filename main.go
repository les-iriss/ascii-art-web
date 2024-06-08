package main

import (
	"fmt"
	"net/http"

	controller "ascii-art-web/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.RootHandler)
	fmt.Println("server running in 'http://localhost:8000'")
	http.ListenAndServe(":8000", mux)
}
