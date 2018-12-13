/* File Name: beer.go
 * Authors: Will Fraisl and Max McKee
 * Description:
 * Usage: Place entire beer folder in $GOPATH/src
 *		  Build and install to $GOBIN with 'go install beer'
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
	// Connect to database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/beer_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	console := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------------------------")
	fmt.Println("Hello and Welcome to the Beer Review Database!!")
	fmt.Println("-----------------------------------------------")

	// Get user type selection
	var sel int
	for true {
		fmt.Print("Are you a (1)brewer, (2)vendor, or (3)rater: ")
		userStr, _ := console.ReadString('\n')
		userStr = strings.TrimSuffix(userStr, "\n")
		sel, _ = strconv.Atoi(userStr)
		if sel > 0 && sel < 4 {
			break
		}
	}

	// Prompt user to log in for their user type
	user := User{}
	var isLoggedIn bool
	for true {
		if isLoggedIn, err = user.Login(db, sel); isLoggedIn {
			break
		}
		if err != nil {
			fmt.Println("There was an error when trying to log in")
			continue
		}
		fmt.Println("Invalid username or password")
	}

	fmt.Printf("Welcome, %s!\n", user.name)

	// Loop to accept user commands
	for true {
		fmt.Print(">")
		input, _ := console.ReadString('\n')
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
				stockBeer(db, string(user.name))
			} else {
				fmt.Println("Only vendor has access to that command")
			}
		case "unstock":
			if user.userType == vendor {
				unstockBeer(db, string(user.name))
			} else {
				fmt.Println("Only vendor has access to that command")
			}
		default:
			fmt.Println("Not a valid command")
		}
	}
}

// Finds all beers that match a name and their brewery, along with vendors that stock it.
func findBeer(db *sql.DB, location string, beerName string) {
	// Check for wildcard
	var request string
	if beerName == "*" {
		request = "SELECT * FROM beer ORDER BY brewery;"
	} else {
		request = "SELECT * FROM beer WHERE name = '" + beerName + "';"
	}
	// Query for all matching beers
	rows, err := db.Query(request)
	if err != nil {
		fmt.Println("Internal error finding beers")
	}
	defer rows.Close()

	// Gets the stock of local vendors for a specific beer
	localStock, err := db.Prepare("SELECT v.name, IFNULL(i.quantity, 0) " +
		"FROM beer b CROSS JOIN vendor v " +
		"LEFT JOIN inventory i ON v.vid = i.vid AND b.name = i.beername AND b.brewery = i.brewery " +
		"WHERE v.location = ? " +
		"AND b.brewery = ? " +
		"AND b.name = ?;")
	if err != nil {
		fmt.Println("There was a problem getting inventory information")
	}

	var name string
	var brewery string
	var abv float32
	var ibu int
	var vendorName string
	var quantity int
	for rows.Next() {
		// Get next beer and print
		err := rows.Scan(&name, &brewery, &abv, &ibu)
		if err != nil {
			continue
		}
		fmt.Printf("%s:  %-35s %.1f ABV \t %d IBU\n", brewery, name, abv, ibu)

		// Query for the local stock of this beer
		localVendors, err := localStock.Query(location, brewery, name)
		defer localVendors.Close()
		if err != nil {
			fmt.Println("    Could not get store information")
			continue
		}

		// Print each vendor and associated stock
		hasVendor := false
		for localVendors.Next() {
			hasVendor = true
			err = localVendors.Scan(&vendorName, &quantity)
			if err != nil {
				fmt.Println("    Could not get store information")
				continue
			}
			fmt.Printf("    %-10s %d in stock\n", vendorName, quantity)
		}
		if !hasVendor {
			fmt.Println("    No vendors found in " + location)
		}
	}
}
