package models

import (
	"time"
)

// Base ...
type Base struct {
	ID      uint64    `db:"id" json:"id"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}
