package controller

import (
	"strings"
	"html/template"
	"net/http"
	fs "ascii-art-web/pkg/fs"
)

type Data struct {
	Text string
	Banner string
	Result string
}
var banners = map[string]bool{
	"shadow": true,
	"standard": true,
	"thinkertoy": true,
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	if r.URL.Path == "/ascii-art" {
		if r.Method == "POST" {
			err_form := r.ParseForm()
			if err_form != nil {
				http.Error(w,err_form.Error(),http.StatusBadRequest)
				return
			}
			data.Text = strings.ReplaceAll(r.FormValue("text"),"\r\n","\\n")
			data.Banner = r.FormValue("banner")
			if !banners[data.Banner] {
				data.Result = "Bad request banner not found! Error 400"
			} else {
				data.Result = fs.AsciiArtFs(data.Text, data.Banner)
			}
		}
	}
	tmpl, err := template.ParseFiles("views/base.html","views/form.html")
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
