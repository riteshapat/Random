package handlers

import (
	"net/http"

	"github.com/riteshapat/go-course/pkg/renders"
)

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.html")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.page.html")
}
