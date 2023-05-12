package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"newSite/utils"
)

var DB *sql.DB

func ConnectDB() {
	var e error
	DB, e = sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=1234 dbname=siteOrbitel sslmode=disable")
	if e != nil {
		fmt.Println(e)
		return
	}

	e = DB.Ping()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	errors := make([]string, 0)

	//errors = append(errors, prepareRequest()...)
	errors = append(errors, prepareUser()...)

	if len(errors) > 0 {
		for _, i := range errors {
			utils.Logger.Println(i)
		}
	}

	LoadSession(sessionMap)
}
