package route

import (
	"net/http"
	"casual-talk/data"
	"casual-talk/utils"
	"html/template"
)

// GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {

}

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := utils.Session(writer, request)
		publicTmplFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		privateTmplFiles := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(privateTmplFiles...))
		} else {
			templates = template.Must(template.ParseFiles(publicTmplFiles...))
		}
		templates.ExecuteTemplate(writer, "layout", threads)
	}
}