/* File Name: init_database.go
 * Authors: Will Fraisl and Max McKee
 * Description: clears beer_database and fills it with example values
 * Usage: go build init_database.go
 * 		  ./init_database
 */

package main

import (
	"database/sql"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/beer_database")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

	file, err := ioutil.ReadFile("init_database.sql")

	// error handling for reading sql file
	if err != nil {
		panic(err.Error())
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			panic(err.Error())
		}
	}
}
