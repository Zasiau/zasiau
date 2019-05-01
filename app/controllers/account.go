package controllers

import (
	"log"
	"net/http"

	"github.com/dongri/candle/app/helpers"
	"github.com/dongri/candle/app/middlewares/render"
	"github.com/dongri/candle/app/services"
	"github.com/mholt/binding"
)

// AccountForm ...
type AccountForm struct {
	Email    string
	Password string
}

// FieldMap ...
func (s *AccountForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
		&s.Password: binding.Field{
			Form:     "password",
			Required: true,
		},
	}
}

// SignupView ...
func SignupView(w http.ResponseWriter, r *http.Request) {
	const view = "account/signup.html"
	output := map[string]interface{}{}
	render.HTML(w, r, view, output)
}

// SignupAction ...
func SignupAction(w http.ResponseWriter, r *http.Request) {
	form := new(AccountForm)
	if err := binding.Bind(r, form); err != nil {
		log.Println(err)
		return
	}
	email := form.Email
	password := form.Password
	user, err := services.AccountSignup(r, email, password)
	if err != nil {
		log.Println(err)
		return
	}
	helpers.SetLoggedInUserID(w, r, user.ID)
	http.Redirect(w, r, "/", http.StatusFound)
}

// SigninView ...
func SigninView(w http.ResponseWriter, r *http.Request) {
	const view = "account/signin.html"
	output := map[string]interface{}{}
	render.HTML(w, r, view, output)
}

// SigninAction ...
func SigninAction(w http.ResponseWriter, r *http.Request) {
	form := new(AccountForm)
	if err := binding.Bind(r, form); err != nil {
		log.Println(err)
		return
	}
	email := form.Email
	password := form.Password
	user, err := services.AccountSignin(r, email, password)
	if err != nil {
		log.Println(err)
		return
	}
	helpers.SetLoggedInUserID(w, r, user.ID)
	http.Redirect(w, r, "/", http.StatusFound)
}

// LogoutAction ...
func LogoutAction(w http.ResponseWriter, r *http.Request) {
	helpers.ClearLoggedInUserID(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
