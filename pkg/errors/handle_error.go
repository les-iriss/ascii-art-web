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

// HandleError sends an HTTP response with the specified error code and renders an error page using the provided error details.
func HandleError(w http.ResponseWriter, r *http.Request, errType Error) {
	w.WriteHeader(errType.Code)
	tmpl, err := template.ParseFiles("views/errors/errors.html")
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	// err = tmpl.Execute(w, error)
	err = tmpl.ExecuteTemplate(w, "error", errType)
	if err != nil {
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
}
