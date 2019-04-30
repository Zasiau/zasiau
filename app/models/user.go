package models

import (
	"fmt"
	"time"

	gorp "gopkg.in/gorp.v1"
)

// User ...
type User struct {
	Base
	UserName string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

// const ...
const (
	UserColumns = "id, username, email, password, created, updated"
)

// Insert ...
func (s *User) Insert(exec gorp.SqlExecutor) error {
	s.Created = time.Now().UTC()
	s.Updated = time.Now().UTC()
	return exec.Insert(s)
}

// Update ...
func (s *User) Update(exec gorp.SqlExecutor) error {
	s.Updated = time.Now().UTC()
	_, err := exec.Update(s)
	return err
}

// UserFindByID ...
func UserFindByID(exec gorp.SqlExecutor, ID uint64) (*User, error) {
	user := new(User)
	query := fmt.Sprintf("SELECT %s FROM users WHERE id = $1", UserColumns)
	if err := exec.SelectOne(&user, query, ID); err != nil {
		return nil, err
	}
	return user, nil
}

// UserFindByEmail ...
func UserFindByEmail(exec gorp.SqlExecutor, email string) (*User, error) {
	user := new(User)
	query := fmt.Sprintf("SELECT %s FROM users WHERE email = $1", UserColumns)
	if err := exec.SelectOne(&user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}
