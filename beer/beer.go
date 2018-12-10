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
	for true{
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
	for true{
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

		switch input {
		case "exit":
			os.Exit(0)
		case "daily":
		case "weekly":
		case "monthly":
		case "yearly":
		case "find":
			// TODO: accept with argument
		case "rate":
			if user.userType == rater {
				rateBeer(db)
			} else {
				fmt.Println("Only rater has access to that command")
			}
		case "add":
			if user.userType == brewer{
				// TODO: have brewery name here aka username
				addBeer(db, "breweryName")
			} else {
				fmt.Println("Only brewery has access to that command")
			}
		case "stock":
			if user.userType == vendor {
				stockBeer(db)
			} else {
				fmt.Println("Only vendor has access to that command")
			}
		case "remove":
			if user.userType == vendor {
				removeBeer(db)
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
func findBeer(db *sql.DB, beerName string) {
	request := "SELECT * FROM beer WHERE name = '" + beerName + "'"
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
func rateBeer(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := reader.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := reader.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Stars (1-5): ")
	stars, _ := reader.ReadString('\n')
	stars = strings.TrimSuffix(stars, "\n")
	fmt.Print("Description (120 characters): ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSuffix(desc, "\n")

	request := "INSERT INTO rating (beer,brewery,stars,description, date) VALUES ('" + beerName + "','" + brewery + "'," + stars + ",'" + desc + "',NOW())"
	_, err := db.Exec(request)
	if err != nil {
		panic(err.Error())
	}
}

// Prompts user for details of new beer they want to add
func addBeer(db *sql.DB, brewery string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Beer Name: ")
	beerName, _ := reader.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Beer Type: ")
	beerType, _ := reader.ReadString('\n')
	beerType = strings.TrimSuffix(beerType, "\n")
	fmt.Print("ABV: ")
	abv, _ := reader.ReadString('\n')
	abv = strings.TrimSuffix(abv, "\n")
	fmt.Print("IBU: ")
	ibu, _ := reader.ReadString('\n')
	ibu = strings.TrimSuffix(ibu, "\n")

	request := "INSERT INTO beer VALUES ('" + beerName + "','" + brewery + "'," + abv + "," + ibu + ")"
	_, err := db.Exec(request)
	if err != nil {
		panic(err.Error())
	}
}

// Prompts user for beer they want to add to their stock
func stockBeer(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := reader.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := reader.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Quantity: ")
	quantity, _ := reader.ReadString('\n')
	quantity = strings.TrimSuffix(quantity, "\n")

	// TODO: stock table in db
}

// Prompts user for beer they want to remove from their stock
func removeBeer(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := reader.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := reader.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")

	// TODO: stock table in db
}