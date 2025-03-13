package graph

import (
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	return pg.Connect(opt)
}
