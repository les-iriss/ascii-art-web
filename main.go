package main

import (
	controller "ascii-art-web/controllers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.RootHandler)
	fmt.Println("server running in 'http://localhost:8000'")
	http.ListenAndServe(":8000", mux)
}
