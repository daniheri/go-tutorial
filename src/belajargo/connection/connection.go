package connection

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var db, err = sql.Open("mysql", "root:@/test")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}
