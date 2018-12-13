/* File Name: rankings.go
 * Authors: Will Fraisl and Max McKee
 * Description: Functions to handle finding the top rated
 *				beers by day, week, month, and year
 */

package main

import (
	"database/sql"
	"fmt"
	"time"
)

// given the database and time frame, return the top beers
func topBeer(db *sql.DB, timeFrame string) {
	currTime := time.Now()
	var queryTime time.Time
	switch timeFrame {
	case "day":
		queryTime = currTime.AddDate(0, 0, -1)
	case "week":
		queryTime = currTime.AddDate(0, 0, -7)
	case "month":
		queryTime = currTime.AddDate(0, -1, 0)
	default: //year
		queryTime = currTime.AddDate(-1, 0, 0)
	}
	queryTimeStr := queryTime.Format("2006-02-01")

	request := "SELECT b.name, AVG(r.stars) "
	request += "FROM beer b JOIN rating r ON (b.name = r.beername) "
	request += "WHERE date >= " + queryTimeStr
	request += " GROUP BY b.name "
	request += "HAVING AVG(r.stars) >= ("
	request += "SELECT AVG(r.stars) "
	request += "FROM beer b JOIN rating r ON (b.name = r.beername) "
	request += "WHERE date >= " + queryTimeStr
	request += " GROUP BY b.name)"
	rows, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var beerName string
	//var brewery string
	var avgRating float32
	for rows.Next() {
		err := rows.Scan(&beerName, &avgRating)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Beer:", beerName, "| Average Rating:", avgRating)
	}
}
