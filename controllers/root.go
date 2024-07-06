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

func executeTemplate(w http.ResponseWriter, r *http.Request, data interface{}) {
	templates, err := template.ParseFiles("views/base.html", "views/form.html")
	if err != nil {
		error.HandleError(w, r, error.Error{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	if err := templates.ExecuteTemplate(w, "base", data); err != nil {
		error.HandleError(w, r, error.Error{Code: http.StatusInternalServerError, Message: err.Error()})
	}
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/ascii-art"{
			w.WriteHeader(404)
			// add something according to the code
			return
		}
	if r.Method != http.MethodPost {
		error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		return
	}
	var data Data
	data.Text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\\n")
	data.Banner = r.FormValue("banner")
	if data.Text == "" {
		error.HandleError(w, r, error.Error{Code: 400, Message: "Bad request empty text! "})
		return
	}
	if len(data.Text) > 250 {
		error.HandleError(w, r, error.Error{Code: 400, Message: "Bad request! You exceeded the length limit."})
		return

	}
	if !banners[data.Banner] {
		error.HandleError(w, r, error.Error{Code: 404, Message: "Bad request banner not found! "})
		return
	}

	data.Result = fs.AsciiArtFs(data.Text, data.Banner)
	executeTemplate(w, r, data)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.WriteHeader(404)
		// here handle it according to the code
		return
	}
	if r.Method != http.MethodGet {
		error.HandleError(w, r, error.Error{Code: 405, Message: "Method not allowed!"})
		return
	}
	executeTemplate(w, r, nil)
}
