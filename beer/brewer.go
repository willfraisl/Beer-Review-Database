package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Prompts user for details of new beer they want to add
func addBeer(db *sql.DB, brewery string) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")
	fmt.Print("Beer Type: ")
	beerType, _ := console.ReadString('\n')
	beerType = strings.TrimSuffix(beerType, "\n")
	fmt.Print("ABV: ")
	abv, _ := console.ReadString('\n')
	abv = strings.TrimSuffix(abv, "\n")
	fmt.Print("IBU: ")
	ibu, _ := console.ReadString('\n')
	ibu = strings.TrimSuffix(ibu, "\n")

	_, err := db.Exec("INSERT INTO beer VALUES(?,?,?,?)", beerName, brewery, abv, ibu)
	if err != nil {
		fmt.Println("Could not add beer. Check that it doesn't exist.")
	}
}

func removeBeer(db *sql.DB, brewery string) {
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Beer Name: ")
	beerName, _ := console.ReadString('\n')
	beerName = strings.TrimSuffix(beerName, "\n")

	_, err := db.Exec("DELETE FROM beer WHERE name = ?", beerName)
	if err != nil {
		fmt.Println("Beer not deleted. Check that it actually exists for your brewery.")
	}
}
