package adapters

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func InitializePostgresql(dbUri string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbUri)
	if err != nil {
		fmt.Println("error connecting to postgresql: ", err)
		os.Exit(1)
	}

	return db
}