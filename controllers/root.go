package controller

import (
	"html/template"
	"net/http"
	"strings"

	errors "ascii-art-web/pkg/errors"
	fs "ascii-art-web/pkg/fs"
)

// Data represents the structure for data used in template rendering.
type Data struct {
	Text   string // User input text to be transformed into ASCII art.
	Banner string // The chosen ASCII art font style.
	Result string // The generated ASCII art result.
}

// banners is a map of available ASCII art font styles.
var banners = map[string]bool{
	"shadow":     true,
	"standard":   true,
	"thinkertoy": true,
}

// executeTemplate compiles and executes the specified templates with provided data.
func executeTemplate(w http.ResponseWriter, r *http.Request, data interface{}) {
	templates, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		errors.HandleError(w, r, errors.Error{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		errors.HandleError(w, r, errors.Error{Code: http.StatusInternalServerError, Message: err.Error()})
	}
}

// PostRequest handles POST requests, processes user input, and renders the ASCII art result.
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

	data.Result = fs.AsciiArtFs(data.Text, data.Banner)
	executeTemplate(w, r, data)
}

// GetRequest handles GET requests and renders the initial form page.
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