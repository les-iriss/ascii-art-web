package errors

import (
	"fmt"
	"html/template"
	"net/http"
)

type Error struct {
	Message string
	Code    int
}

func HandleError(w http.ResponseWriter, r *http.Request, error Error) {
	tmpl, err := template.ParseFiles("views/errors/errors.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(error.Code)
}
