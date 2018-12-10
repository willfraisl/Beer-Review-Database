/* File Name: user.go
 * Authors: Will Fraisl and Max McKee
 * Description: Generalizes logged in users to maintain 
 * 				control of what users can and can't do
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"database/sql"
	"crypto/sha256"
	"golang.org/x/crypto/ssh/terminal"
)

// User type enumeration
type UserType string
const(
	brewer UserType = "brewery"
	vendor UserType = "vendor"
	rater  UserType = "rater"
)

// Generalized user struct for breweries, raters, and vendors
type User struct{
	userType UserType
	name string
}

// Accepts a database connection and a user type where 
// 1 -> brewery, 2 -> vendor, 3 -> rater
// Returns a boolean indicating success and any errors
func (u *User) Login(db *sql.DB, userType int) (bool, error){
	// Prompt for username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, err := reader.ReadString('\n')
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
	if err != nil{
		return false, err
	}
	defer rows.Close()
	
	// Check if a response is sent for the user with the given type
	if rows.Next() {
		var pHash string
		// Fetch data from the first row found
		if err = rows.Scan(&pHash); err != nil{
			return false, err
		} else if pHash == fmt.Sprintf("%x", h.Sum(nil)) { // Sucessful login
			u.name = username
			return true, nil
		} else { // Failed login
			return false, err
		}
	}
	
	return false, nil
}