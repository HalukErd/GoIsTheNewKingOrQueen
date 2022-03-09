package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

const driverName = "mysql"
const user = "admin"
const password = "password"
const endpoint = "database-1.cj5hqjeb7vmg.us-east-1.rds.amazonaws.com"
const port = "3306"
const dbname = "mydb"

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, endpoint, port, dbname))
	checkErr()
	defer db.Close()
	err = db.Ping()
	checkErr()
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	checkErr()
}

func checkErr() {
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
	checkErr()
}
