package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Prompts user for beer they want to add to their stock
func stockBeer(db *sql.DB) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := console.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Quantity: ")
	quantity, _ := console.ReadString('\n')
	quantity = strings.TrimSuffix(quantity, "\n")

	// TODO: stock table in db
	// test push
}

// Prompts user for beer they want to remove from their stock
func unstockBeer(db *sql.DB) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := console.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")

	// TODO: stock table in db
}
