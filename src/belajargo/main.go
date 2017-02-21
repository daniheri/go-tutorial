package main

import (
	"net/http"
	"belajargo/model"
	"fmt"
	"belajargo/query"
	"belajargo/connection"
	"os"
	"encoding/json"
)
func main() {
	query.SqlQuery()
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func user(w http.ResponseWriter, r *http.Request) {
	var db = connection.Connect()
	defer db.Close()

	var rows, err = db.Query("select id, name, grade, age from tb_student")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close()

	var result []model.Student

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		for rows.Next() {
			var each = model.Student{}
			var err = rows.Scan(&each.Id, &each.Name, &each.Grade, &each.Age)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}

			var result, error = json.Marshal(result)

			if error != nil {
				http.Error(w, error.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}