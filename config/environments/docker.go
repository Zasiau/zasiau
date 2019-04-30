package environments

import "os"

type docker struct{}

func (c docker) PostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}
