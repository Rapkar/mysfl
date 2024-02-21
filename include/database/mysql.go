package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect() bool {
	db, err := sql.Open("mysql", "root:<0311121314>@tcp(127.0.0.1:3306)/GOLANG")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return true
}
