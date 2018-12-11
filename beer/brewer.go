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
