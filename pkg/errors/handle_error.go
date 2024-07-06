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

// error is a reserved name for golang , you can not use it as a variable
func HandleError(w http.ResponseWriter, r *http.Request, er Error) {
	tmpl, err := template.ParseFiles("views/errors/errors.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
	// err = tmpl.Execute(w, error)
	w.WriteHeader(er.Code)
	err = tmpl.ExecuteTemplate(w, "error", er)
	// mybe we have to use log, because using fmt with w.Write is not a good idea
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		fmt.Println(err.Error(), http.StatusInternalServerError)
		return
	}
}
