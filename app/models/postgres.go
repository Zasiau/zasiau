package models

import gorp "gopkg.in/gorp.v1"

// AddTableWithName ...
func AddTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	dbmap.AddTableWithName(Product{}, "products").SetKeys(true, "ID")
}
