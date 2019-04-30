package environments

import (
	"fmt"
)

// Env ...
type Env string

// Const ...
const (
	Docker      Env = "docker"
	Development Env = "development"
	Production  Env = "production"
)

// CurrentEnv ...
var CurrentEnv Env

func (e Env) String() string {
	return string(e)
}

// EnvConfig ...
type EnvConfig interface {
	PostgresURI() string
}

// CurrentConfig ...
var CurrentConfig EnvConfig

// SetEnvironment ..
func SetEnvironment(e Env) error {
	var configInstance EnvConfig
	switch e {
	case Production:
		configInstance = &production{}
	case Development:
		configInstance = &development{}
	case Docker:
		configInstance = &docker{}
	}
	if configInstance == nil {
		return fmt.Errorf("No configuration found for %s", e)
	}

	CurrentConfig = configInstance
	CurrentEnv = e
	return nil
}
