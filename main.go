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


// you can not use the name error or as a variable 
// you can use r.ContentLenght instead of len() but it is fine
// you should write to the header using w.Write(status code) 
// and not just return a message to the user 
// in the fs file you should return an error alogside the ascii art fs 
// if there is an error you should use return ascii, error , and not just ignore it
// you should respond with the proper status code and not use them randomly


// things i didn't fix :
// return erros in the fs package
