package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates inserts the html templates into the templates variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate renders an html page on screen
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
