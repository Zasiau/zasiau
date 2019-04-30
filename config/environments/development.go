package environments

import "os"

type development struct{}

func (c development) PostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}
