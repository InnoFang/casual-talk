package utils

import (
	"net/http"
	"fmt"
	"html/template"
	"strings"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("casual-talk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// parse HTML templates
// pass in a list of file names, and get a template
func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func GenerateHTML(writer http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// convenience function to redirect to the error message page
func ErrorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// log
func Info(args ...interface{}) {
	logger.SetPrefix("[INFO] ")
	logger.Println(args...)
}

func Danger(args ...interface{}) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(args...)
}

func Warn(args ...interface{}) {
	logger.SetPrefix("[WARN] ")
	logger.Println(args...)
}