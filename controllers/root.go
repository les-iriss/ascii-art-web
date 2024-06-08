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
		postRequest(w, r, &data)
		getPage(w, &data)
	case "/":
		if r.Method == "GET" {
			getPage(w, &data)
		} else {
			error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		}
	default:
		error.HandleError(w, r, error.Error{Code: 404, Message: "Page not found!"})
	}
}

func postRequest(w http.ResponseWriter, r *http.Request, data *Data) {
	if r.Method == "POST" {
		err_form := r.ParseForm()
		if err_form != nil {
			http.Error(w, err_form.Error(), http.StatusBadRequest)
			return
		}
		data.Text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\\n")
		data.Banner = r.FormValue("banner")
		if !banners[data.Banner] {
			data.Result = "Bad request banner not found! Error 400."
			return
		} else {
			data.Result = fs.AsciiArtFs(data.Text, data.Banner)
		}
	} else {
		error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		return
	}
}

func getPage(w http.ResponseWriter, data *Data) {
	tmpl, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
