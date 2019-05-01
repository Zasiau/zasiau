package render

import (
	"context"
	"net/http"

	"github.com/flosch/pongo2"
	renderLib "github.com/unrolled/render"

	"github.com/dongri/candle/app/helpers"
	"github.com/dongri/candle/config/environments"
)

type contextKey string

// const ...
const (
	renderContextKey contextKey = "render/render_setup"
)

// Builder ...
type Builder struct {
	Env environments.Env
}

// MiddleWare ...
func (b *Builder) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		renderEngine := renderLib.New(renderLib.Options{
			IndentJSON:    (r.Form.Get("pretty") == "true"),
			IsDevelopment: (b.Env == environments.Development),
		})
		r = r.WithContext(context.WithValue(r.Context(), renderContextKey, renderEngine))
		next.ServeHTTP(w, r)
	})
}

// GetRenderer ...
func GetRenderer(r *http.Request) *renderLib.Render {
	return r.Context().Value(renderContextKey).(*renderLib.Render)
}

// JSON ...
func JSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render := GetRenderer(r)
	render.JSON(w, status, data)
}

// HTML ...
func HTML(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	file = "app/views/" + file
	loggedIn := helpers.GetLoggedInUserID(r) != 0
	context := pongo2.Context{
		"logged_in": loggedIn,
	}
	tpl, err := pongo2.DefaultSet.FromFile(file)
	if err != nil {
		panic(err)
	}
	if data != nil {
		for k, v := range data.(map[string]interface{}) {
			context[k] = v
		}
	}
	tpl.ExecuteWriter(context, w)
}
