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
