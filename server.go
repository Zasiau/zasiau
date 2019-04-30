package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"

	"github.com/dongri/gonion/app"
	"github.com/dongri/gonion/app/models"
	"github.com/dongri/gonion/config/environments"
)

const (
	defaultEnv  = environments.Development
	defaultPort = "3000"
)

func main() {
	env := environments.Env(os.Getenv("GO_ENV"))
	if env == "" {
		env = environments.Env(defaultEnv)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	environments.SetEnvironment(env)

	// setup postgres
	db, err := sql.Open("postgres", environments.CurrentConfig.PostgresURI())
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer db.Close()
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	models.AddTableWithName(dbmap)

	// setup session store
	sessionStore := sessions.NewCookieStore([]byte("@xGYEX%TfUCdG8JGzY,DKtE7[kTKegvqWvXb"))

	rooter := app.NewRootHandler(env, dbmap, sessionStore)

	// Launch a server instance.
	fmt.Printf("Server listening on port %s in %s mode.\n", port, env)
	http.ListenAndServe(":"+port, rooter)
}
