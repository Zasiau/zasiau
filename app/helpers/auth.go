package helpers

import (
	"net/http"

	"github.com/dongri/gonion/app/middlewares/session"
)

const (
	sessionKeyUser = "user"
)

// SetLoggedInUserID ...
func SetLoggedInUserID(w http.ResponseWriter, r *http.Request, userID uint64) {
	store := session.GetSessionStore(r)
	session, _ := store.Get(r, sessionKeyUser)
	session.Values["user_id"] = userID
	session.Save(r, w)
}

// GetLoggedInUserID ...
func GetLoggedInUserID(r *http.Request) uint64 {
	store := session.GetSessionStore(r)
	session, _ := store.Get(r, sessionKeyUser)
	userID := session.Values["user_id"]
	if userID != nil {
		return userID.(uint64)
	}
	return uint64(0)
}

// ClearLoggedInUserID ...
func ClearLoggedInUserID(w http.ResponseWriter, r *http.Request) {
	store := session.GetSessionStore(r)
	session, _ := store.Get(r, sessionKeyUser)
	session.Values["user_id"] = uint64(0)
	session.Save(r, w)
}
