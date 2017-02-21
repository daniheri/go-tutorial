package query

import (
	"fmt"
	"os"
	"belajargo/connection"
	"belajargo/model"
)

func SqlQuery() {
	var db = connection.Connect()
	defer db.Close()

	var rows, err = db.Query("select id, name, grade, age from tb_student")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close()

	var result []model.Student

	for rows.Next() {
		var each = model.Student{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Grade, &each.Age)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
