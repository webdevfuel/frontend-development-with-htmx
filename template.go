package main

import "html/template"

func parseTemplates() (*template.Template, error) {
	return template.ParseGlob("templates/*.html")
}
