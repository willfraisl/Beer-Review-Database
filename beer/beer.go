/* File Name: beer.go
 * Authors: Will Fraisl and Max McKee
 * Description:
 * Usage: go build beer.go
 * 		  ./beer
 */

package main

import (
	"bufio"
	"database/sql"
	"fmt"
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

	var sel int
	for true {
		fmt.Print("Are you a (1)brewer, (2)vendor, or (3)rater: ")
		userStr, _ := reader.ReadString('\n')
		userStr = strings.TrimSuffix(userStr, "\n")
		sel, _ = strconv.Atoi(userStr)
		if sel > 0 && sel < 4 {
			break
		}
	}

	// check what type of user

	// connect to database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/beer_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	user := User{}
	var isLoggedIn bool
	for true {
		if isLoggedIn, err = user.Login(db, sel); isLoggedIn {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Invalid username or password")
	}

	fmt.Printf("Welcome, %s!\n", user.name)

	// loop to accept user commands
	for true {
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		cmd := strings.SplitN(input, " ", 2)

		switch cmd[0] {
		case "exit":
			os.Exit(0)
		case "daily":
			topBeer(db, "day")
		case "weekly":
			topBeer(db, "week")
		case "monthly":
			topBeer(db, "month")
		case "yearly":
			topBeer(db, "year")
		case "find":
			// TODO: accept with argument
			findBeer(db, user.location, cmd[1])
		case "rate":
			if user.userType == rater {
				rateBeer(db)
			} else {
				fmt.Println("Only rater has access to that command")
			}
		case "add":
			if user.userType == brewer {
				addBeer(db, string(user.name))
			} else {
				fmt.Println("Only brewery has access to that command")
			}
		case "remove":
			if user.userType == brewer {
				removeBeer(db, string(user.name))
			} else {
				fmt.Println("Only brewery has access to that command")
			}
		case "stock":
			if user.userType == vendor {
				stockBeer(db)
			} else {
				fmt.Println("Only vendor has access to that command")
			}
		case "unstock":
			if user.userType == vendor {
				unstockBeer(db)
			} else {
				fmt.Println("Only vendor has access to that command")
			}
		default:
			fmt.Println("Not a valid command")
			// TODO: print a list of commands
		}
	}
}

// Finds all beers that match a name and their brewery, along with vendors that stock it.
func findBeer(db *sql.DB, location string, beerName string) {
	var request string
	if beerName == "*" {
		request = "SELECT * FROM beer ORDER BY brewery;"
	} else {
		request = "SELECT * FROM beer WHERE name = '" + beerName + "';"
	}
	rows, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Gets the stock of local vendors for a specific beer
	localStock, err := db.Prepare("SELECT v.name, i.quantity " +
		"FROM beer b CROSS JOIN vendor v " +
		"LEFT JOIN inventory i ON v.vid = i.vid AND b.name = i.beername AND b.brewery = i.brewery " +
		"WHERE v.location = ? " +
		"AND b.brewery = ? " +
		"AND b.name = ?;")
	if err != nil {
		panic(err.Error())
	}
	var name string
	var brewery string
	var abv float32
	var ibu int
	var vendorName string
	var quantity int
	for rows.Next() {
		err := rows.Scan(&name, &brewery, &abv, &ibu)
		if err != nil {
			continue
		}
		fmt.Printf("%s:  %-35s %.1f ABV \t %d IBU\n", brewery, name, abv, ibu)

		localVendors, err := localStock.Query(location, brewery, name)
		defer localVendors.Close()
		if err != nil {
			panic(err.Error())
		}

		hasVendor := false
		for localVendors.Next() {
			hasVendor = true
			err = localVendors.Scan(&vendorName, &quantity)
			if err != nil {
				quantity = 0
			}
			fmt.Printf("    %-10s %d in stock\n", vendorName, quantity)
		}
		if !hasVendor {
			fmt.Println("    No vendors found in " + location)
		}
	}
}
