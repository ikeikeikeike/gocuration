package main

import (
	"database/sql"
	"os"
	"testing"

	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestConnDatabase(t *testing.T) {
	driver := "sqlite3"
	connStr := "./antenna_development.db"

	db, err := sql.Open(driver, connStr)
	if err != nil {
		pp.Println("[ERRO]", err.Error(), driver, connStr)
		os.Exit(2)
	}

	println(db)
}
