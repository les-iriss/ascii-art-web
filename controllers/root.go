package controller

import (
	"html/template"
	"net/http"
	"strings"

	errors "ascii-art-web/pkg/errors"
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

func executeTemplate(w http.ResponseWriter, r *http.Request, data interface{}) {
	templates, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		errors.HandleError(w, r, errors.Error{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		errors.HandleError(w, r, errors.Error{Code: http.StatusInternalServerError, Message: err.Error()})
	}
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.HandleError(w, r, errors.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed!"})
		return
	}
	var data Data
	data.Text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\\n")
	data.Banner = r.FormValue("banner")
	if !banners[data.Banner] {
		errors.HandleError(w, r, errors.Error{Code: http.StatusBadRequest, Message: "Bad request invalid banner! "})
		return
	}
	if data.Text == "" {
		errors.HandleError(w, r, errors.Error{Code: http.StatusBadRequest, Message: "Bad request empty text! "})
		return
	}
	if len(data.Text) > 250 {
		errors.HandleError(w, r, errors.Error{Code: http.StatusBadRequest, Message: "Bad request! You exceeded the length limit."})
		return

	}
	if !banners[data.Banner] {
		errors.HandleError(w, r, errors.Error{Code: http.StatusBadRequest, Message: "Bad request banner not found! "})
		return
	}

	data.Result = fs.AsciiArtFs(data.Text, data.Banner)
	executeTemplate(w, r, data)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errors.HandleError(w, r, errors.Error{Code: http.StatusNotFound, Message: "Page not found!"})
		return
	}
	if r.Method != http.MethodGet {
		errors.HandleError(w, r, errors.Error{Code: http.StatusMethodNotAllowed, Message: "Method not allowed!"})
		return
	}
	executeTemplate(w, r, nil)
}
