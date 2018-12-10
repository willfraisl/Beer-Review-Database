/* File Name: user.go
 * Authors: Will Fraisl and Max McKee
 * Description: Generalizes logged in users to maintain 
 * 				control of what users can and can't do
 */

package main

import (
	"bufio"
	"os"
	"fmt"
)

// User type enumeration
type UserType int
const(
	brewer UserType = 1
	vendor UserType = 2
	rater  UserType = 3
)

// Generali
type User struct{
	userType UserType
	name string
}

func (u *User) login(userType int) (bool, error){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')

	switch userType {
	case 1:
		// TODO: check username and pass
		u.userType = brewer
	case 2:
		// TODO: check username and pass
		u.userType = vendor
	case 3:
		// TODO: check username and pass
		u.userType = rater
	default:
	}
	
	fmt.Println("Username is " + username)
	fmt.Println("Password is " + password)
	return false, nil
}