package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func MysqlConnect() bool {
	envFile, _ := godotenv.Read("config.env")
	Name := envFile["DATABASENAME"]
	User := envFile["DATABASEUSER"]
	Pass := envFile["DATABASEPASS"]
	db, err := sql.Open("mysql", User+":<"+Pass+">@tcp(127.0.0.1:3306)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return true
}
