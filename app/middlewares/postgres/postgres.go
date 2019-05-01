package postgres

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	gorp "gopkg.in/gorp.v1"

	"github.com/dongri/candle/config/environments"
)

type contextKey string

const (
	postgresContextKey contextKey = "filters/postgres"
)

// Builder ...
type Builder struct {
	DbMap *gorp.DbMap
	next  http.Handler
}

// MiddleWare ...
func (b *Builder) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), postgresContextKey, b.DbMap))
		next.ServeHTTP(w, r)
	})
}

// GetDbMap ...
func GetDbMap(r *http.Request) *gorp.DbMap {
	dbmap, ok := r.Context().Value(postgresContextKey).(*gorp.DbMap)
	if !ok || dbmap == nil {
		log.Printf("Postgres is not set in the request")
		db, err := sql.Open("postgres", environments.CurrentConfig.PostgresURI())
		if err != nil {
			log.Print(err.Error())
			log.Fatal("Failed to connect to Postgres")
			return nil
		}
		dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	}
	return dbmap
}
