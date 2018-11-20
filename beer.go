/* File Name: beer.go
 * Authors: Will Fraisl and Max McKee
 * Description:
 * Usage: go build beer.go
 * ./beer
 */

package main

import (
    "fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello and welcome to the Beer Review Database!!")
	fmt.Println("Are you a brewer, vendor, or rater?")
	fmt.Println("Enter username: ")
	fmt.Println("Enter password: ")

    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/beer_database")
    
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
    
    // defer the close till after the main function has finished
	defer db.Close()
}