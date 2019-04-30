package session

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

type contextKey string

const (
	sessionStoreContextKey contextKey = "filters/session_store"
)

// Builder ...
type Builder struct {
	SessionStore *sessions.CookieStore
	next         http.Handler
}

// MiddleWare ...
func (b *Builder) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), sessionStoreContextKey, b.SessionStore))
		next.ServeHTTP(w, r)
	})
}

// GetSessionStore ...
func GetSessionStore(r *http.Request) *sessions.CookieStore {
	return r.Context().Value(sessionStoreContextKey).(*sessions.CookieStore)
}
