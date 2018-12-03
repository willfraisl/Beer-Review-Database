/* File Name: beer.go
 * Authors: Will Fraisl and Max McKee
 * Description:
 * Usage: go build beer.go
 * 		  ./beer
 */

package main

import (
    "fmt"
	"database/sql"
	"bufio"
	"os"
	"strconv"
	"strings"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Hello and Welcome to the Beer Review Database!!")
	fmt.Println("-----------------------------------------------")
	
	fmt.Print("Are you a (1)brewer, (2)vendor, or (3)rater: ")
	user_str, _ := reader.ReadString('\n')
	user_str = strings.TrimSuffix(user_str, "\n")
	user, err := strconv.Atoi(user_str)

	switch user {
	case 1:
		fmt.Println("brewer")
	case 2:
		fmt.Println("vendor")
	case 3:
		fmt.Println("rater")
	default:
		fmt.Println("wrong input")
	}

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')

	fmt.Println("Username: " + username)
	fmt.Println("Password: " + password)

    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/beer_database")    
    if err != nil {
        panic(err.Error())
	}	
	defer db.Close()

	top_beer(db, "daily")
	fmt.Println()
	find_beer(db, "Falls Porter")
	fmt.Println()
	rate_beer(db)
	fmt.Println()
}

// given the database and time frame, return the top beers
func top_beer(db *sql.DB, time_frame string) {
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

// Finds all beers that match a name and their brewery, along with vendors that stock it.
func find_beer(db *sql.DB, beer_name string) {
	request := "SELECT * FROM beer WHERE name = '" + beer_name + "'"
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

// Prompts user for rating details and stores their rating. (max 1 rating per beer per day)
func rate_beer(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := reader.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beer_name, _ := reader.ReadString('\n')
	beer_name = strings.TrimSuffix(beer_name, "\n")
	fmt.Print("Stars (1-5): ")
	stars, _ := reader.ReadString('\n')
	stars = strings.TrimSuffix(stars, "\n")
	fmt.Print("Description (120 characters): ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSuffix(desc, "\n")

	request := "INSERT INTO rating (beer,brewery,stars,description, date) VALUES ('"+beer_name+"','"+brewery+"',"+stars+",'"+desc+"',NOW())"
	_, err := db.Exec(request)
	if err != nil {
		panic(err.Error())
	}
}

// Prompts user for details of new beer they want to add
func add_beer(db *sql.DB) {

}

// Prompts user for beer they want to add to their stock
func stock_beer(db *sql.DB) {

}

// Prompts user for beer they want to remove from their stock
func remove_beer(db *sql.DB) {

}