package controllers

import (
	"net/http"

	"github.com/dongri/gonion/app/middlewares/render"
)

// HomeIndexHandler ...
func HomeIndexHandler(w http.ResponseWriter, r *http.Request) {
	const view = "home/index.html"
	output := map[string]interface{}{
		"message": "hello",
	}
	render.HTML(w, r, view, output)
}
