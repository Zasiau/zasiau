package app

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/gorp.v1"

	"github.com/dongri/gonion/app/controllers"
	"github.com/dongri/gonion/app/controllers/product"
	"github.com/dongri/gonion/app/helpers"
	"github.com/dongri/gonion/app/middlewares/postgres"
	"github.com/dongri/gonion/app/middlewares/render"
	"github.com/dongri/gonion/app/middlewares/session"
	"github.com/dongri/gonion/config/environments"
)

var defaultMiddlewares = []func(next http.Handler) http.Handler{
	middleware.Logger,
	middleware.Recoverer,
	middleware.StripSlashes,
}

// NewRootHandler returns a root handler.
func NewRootHandler(env environments.Env, dbmap *gorp.DbMap, sessionStore *sessions.CookieStore) http.Handler {
	r := chi.NewRouter()

	// setup middlewares
	renderBuilder := &render.Builder{Env: env}
	middlewares := append(defaultMiddlewares, renderBuilder.MiddleWare)

	postgresBuilder := &postgres.Builder{DbMap: dbmap}
	middlewares = append(middlewares, postgresBuilder.MiddleWare)

	sessionStoreBuilder := &session.Builder{SessionStore: sessionStore}
	middlewares = append(middlewares, sessionStoreBuilder.MiddleWare)

	r.Use(middlewares...)

	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/stylesheets"))))
	r.Handle("/img/*", http.StripPrefix("/img/", http.FileServer(http.Dir("./public/images"))))
	r.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/javascripts"))))

	r.Get("/", controllers.HomeIndexHandler)

	r.Get("/signup", controllers.SignupView)
	r.Post("/signup", controllers.SignupAction)
	r.Get("/signin", controllers.SigninView)
	r.Post("/signin", controllers.SigninAction)
	r.Get("/logout", controllers.LogoutAction)

	r.Mount("/products", productRouter())

	return r
}

// /account/*
func productRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(auth)
	r.Get("/", product.Index)
	r.Get("/{id}", product.Show)
	r.Get("/new", product.New)
	r.Get("/{id}/edit", product.Edit)
	r.Post("/create", product.Create)
	r.Post("/{id}", product.Update)
	r.Get("/{id}/delete", product.Destroy)
	return r
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helpers.GetLoggedInUserID(r) == uint64(0) {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
