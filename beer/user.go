/* File Name: user.go
 * Authors: Will Fraisl and Max McKee
 * Description: Generalizes logged in users to maintain
 * 				control of what users can and can't do
 */

package main

import (
	"bufio"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// User type enumeration
type UserType string

const (
	brewer UserType = "brewery"
	vendor UserType = "vendor"
	rater  UserType = "rater"
)

// Generalized user struct for breweries, raters, and vendors
type User struct {
	userType UserType
	name     string
}

// Accepts a database connection and a user type where
// 1 -> brewery, 2 -> vendor, 3 -> rater
// Returns a boolean indicating success and any errors
func (u *User) Login(db *sql.DB, userType int) (bool, error) {
	// Prompt for username
	console := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, err := console.ReadString('\n')
	if err != nil {
		return false, err
	}
	username = strings.TrimSpace(username)

	// Prompt for password, don't echo to terminal
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return false, err
	}
	fmt.Println()

	// Determine and set user type
	switch userType {
	case 1:
		u.userType = brewer
	case 2:
		u.userType = vendor
	case 3:
		u.userType = rater
	default:
	}

	// Hash password
	h := sha256.New()
	h.Write([]byte(password))

	// Run query to fetch password hash from database for given users
	qString := ("SELECT password FROM " + string(u.userType) +
		" WHERE name = '" + username + "';")
	rows, err := db.Query(qString)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Check if a response is sent for the user with the given type
	if rows.Next() {
		var pHash string
		// Fetch data from the first row found
		if err = rows.Scan(&pHash); err != nil {
			return false, err
		} else if pHash == fmt.Sprintf("%x", h.Sum(nil)) { // Sucessful login
			u.name = username
			return true, nil
		} else { // Failed login
			return false, err
		}
	} else { // No user exists
		// Loop for valid input
		for true {
			// Prompt user if they want to create the user
			fmt.Printf("The %s does not exist yet.\n", string(u.userType))
			fmt.Print("Create a new account with the given password? (y/n): ")
			response, err := console.ReadString('\n')
			if err != nil {
				return false, err
			}
			response = strings.TrimSpace(response)

			if response == "n" {
				fmt.Println("Okay, reprompting for credentials...")
				return false, nil
			} else if response == "y" {
				break
			}
		}

		// Construct update statement
		u.name = username
		updateString := ("INSERT INTO " + string(u.userType) + " VALUES ('" +
			u.name + "', '" + fmt.Sprintf("%x", h.Sum(nil)) + "'")

		// New brewer and vendor requires location
		if u.userType == brewer || u.userType == vendor {
			fmt.Print("Please enter your location of business: ")
			location, err := console.ReadString('\n')
			if err != nil {
				return false, err
			}
			location = strings.TrimSpace(location)
			updateString += ",'" + location + "'"
		}

		updateString += ");"

		// Add user to the database
		_, err := db.Exec(updateString)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
