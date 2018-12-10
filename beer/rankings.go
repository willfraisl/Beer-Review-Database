/* File Name: rankings.go
 * Authors: Will Fraisl and Max McKee
 * Description: Functions to handle finding the top rated
 *				beers by day, week, month, and year
 */
package main

import (
	"fmt"
	"database/sql"
)

// given the database and time frame, return the top beers
func topBeer(db *sql.DB, timeFrame string) {
	request := "SELECT * FROM beer"
	rows, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var name string
	var brewery string
	var abv float32
	var ibu int
	for rows.Next() {
		err := rows.Scan(&name, &brewery, &abv, &ibu)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(name, brewery, abv, ibu)
	}
}