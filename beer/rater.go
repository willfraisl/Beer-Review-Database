package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Prompts user for rating details and stores their rating. (max 1 rating per beer per day)
func rateBeer(db *sql.DB) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := console.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Stars (1-5): ")
	stars, _ := console.ReadString('\n')
	stars = strings.TrimSuffix(stars, "\n")
	fmt.Print("Description (120 characters): ")
	desc, _ := console.ReadString('\n')
	desc = strings.TrimSuffix(desc, "\n")

	// update := "INSERT INTO rating (beer,brewery,stars,description, date) VALUES ('" + beerName + "','" + brewery + "'," + stars + ",'" + desc + "',NOW())"
	_, err := db.Exec("INSERT INTO rating (beer,brewery,stars,description, date) VALUES (?,?,?,?,NOW())",
		beerName, brewery, stars, desc)
	if err != nil {
		panic(err) //fmt.Println("Cannot add rating. (Maybe the brewery hasn't added the beer?)")
	}
}
