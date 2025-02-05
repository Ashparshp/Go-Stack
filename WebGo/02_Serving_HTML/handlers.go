package main

import (
	"net/http"
)

func HomePage (w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl")
}
func AboutPage (w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.tmpl")
}