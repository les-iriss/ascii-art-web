package controller

import (
	"html/template"
	"net/http"
	"strings"

	error "ascii-art-web/pkg/errors"
	fs "ascii-art-web/pkg/fs"
)

type Data struct {
	Text   string
	Banner string
	Result string
}

var banners = map[string]bool{
	"shadow":     true,
	"standard":   true,
	"thinkertoy": true,
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	switch r.URL.Path {
	case "/ascii-art":
		if r.Method == "POST" {
			postRequest(w, r, &data)
		} else {
			error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		}
		// getPage(w, &data)
	case "/":
		if r.Method == "GET" {
			getPage(w, r, &data)
		} else {
			error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		}
	default:
		error.HandleError(w, r, error.Error{Code: 404, Message: "Page not found!"})
	}
	w.WriteHeader(http.StatusOK)

}

func postRequest(w http.ResponseWriter, r *http.Request, data *Data) {
	tmpl, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		error.HandleError(w, r, error.Error{Code: 400, Message: "Bad request banner not found! "})

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.Text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\\n")
	data.Banner = r.FormValue("banner")
	if !banners[data.Banner] {
		// data.Result = "Bad request banner not found! Error 400."
		error.HandleError(w, r, error.Error{Code: 400, Message: "Bad request banner not found! "})
		http.Error(w, "", http.StatusBadRequest)
		return
	} else {
		data.Result = fs.AsciiArtFs(data.Text, data.Banner)
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPage(w http.ResponseWriter, r *http.Request, data *Data) {
	tmpl, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		error.HandleError(w, r, error.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		error.HandleError(w, r, error.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}
