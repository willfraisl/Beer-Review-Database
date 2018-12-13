package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Prompts user for beer they want to add to their stock
func stockBeer(db *sql.DB, vendor string) {
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

	// TODO: fix so always returns 1 row
	request := "SELECT vid FROM vendor WHERE name = '" + vendor + "'"
	rows, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var vendorid string
	var vid string
	for rows.Next() {
		err := rows.Scan(&vid)
		if err != nil {
			panic(err.Error())
		}
		vendorid = vid
	}

	_, err = db.Exec("INSERT INTO inventory VALUES(?,?,?,?)", vendorid, beerName, brewery, quantity)
	if err != nil {
		fmt.Println("Could not add stock. Check that stock doesn't already exist.")
	} else {
		fmt.Println("Successfully added stock.")
	}
}

// Prompts user for beer they want to remove from their stock
func unstockBeer(db *sql.DB, vendor string) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Brewery: ")
	brewery, _ := console.ReadString('\n')
	brewery = strings.TrimSuffix(brewery, "\n")
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")

	// TODO: fix so always returns 1 row
	request := "SELECT vid FROM vendor WHERE name = '" + vendor + "'"
	rows, err := db.Query(request)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var vendorid string
	var vid string
	for rows.Next() {
		err := rows.Scan(&vid)
		if err != nil {
			panic(err.Error())
		}
		vendorid = vid
	}

	_, err = db.Exec("DELETE FROM inventory WHERE vid=? AND beer=? AND brewery=?", vendorid, beerName, brewery)
	if err != nil {
		fmt.Println("Inventory not deleted. Check that it actually exists in your stock.")
	} else {
		fmt.Println("Stock deleted successfully.")
	}
}
