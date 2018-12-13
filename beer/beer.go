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
			findBeer(db, cmd[1])
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
			// TODO: print a list of commands
		}
	}
}

// Finds all beers that match a name and their brewery, along with vendors that stock it.
func findBeer(db *sql.DB, beerName string) {
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

	var name string
	var brewery string
	var abv float32
	var ibu int
	for rows.Next() {
		err := rows.Scan(&name, &brewery, &abv, &ibu)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%s:  %-35s %.1f ABV \t %d IBU\n", brewery, name, abv, ibu)
		// TODO: Left join with inventory table to see who stocks each beer
	}
}
