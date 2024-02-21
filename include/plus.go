package include

import (
	"fmt"
	"include/include/database"
)

func BootCore() {

	if database.MysqlConnect() {
		fmt.Println("database succes!")
	}
}
