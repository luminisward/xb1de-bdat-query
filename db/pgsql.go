package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"xb1de-bdat-query/config"
)

var db *sqlx.DB

func GetPgsql() *sqlx.DB {
	if db == nil {
		dbconf := config.GetConfig().DbConfig
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.Dbname)

		var err error
		db, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return db
}
