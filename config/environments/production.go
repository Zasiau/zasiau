package environments

import "os"

type production struct{}

func (c production) PostgresURI() string {
	return os.Getenv("DATABASE_URL")
}
