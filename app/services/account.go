package services

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dongri/gonion/app/middlewares/postgres"
	"github.com/dongri/gonion/app/models"
	"golang.org/x/crypto/bcrypt"
)

// AccountSignup ...
func AccountSignup(r *http.Request, email, password string) (*models.User, error) {
	dbmap := postgres.GetDbMap(r)
	user, err := models.UserFindByEmail(dbmap, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Singup Error")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Singup Error")
	}
	user = new(models.User)
	user.Email = email
	user.Password = string(hashedPassword)
	if err := user.Insert(dbmap); err != nil {
		return nil, err
	}
	return user, nil
}

// AccountSignin ...
func AccountSignin(r *http.Request, email, password string) (*models.User, error) {
	dbmap := postgres.GetDbMap(r)
	user, err := models.UserFindByEmail(dbmap, email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Password Error")
	}
	return user, nil
}
