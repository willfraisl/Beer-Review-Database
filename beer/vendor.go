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
