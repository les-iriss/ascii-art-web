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

		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	// err = tmpl.Execute(w, error)
	err = tmpl.ExecuteTemplate(w, "error", error)
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
}
