/*
1. buat login
2. require username & password
3. username > string
4. password > **** (len)

if username || password == true
username dan password benar
else
username/password salah
*/
package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	defer catchErr()

	// var password string
	var username string
	fmt.Println("Input Username:")
	fmt.Scanln(&username)
	fmt.Println("Input Password:")
	// fmt.Scanln(&password)
	password, _ := gopass.GetPasswdMasked()

	if valid, err := validateAccount(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validateAccount(username string, password []byte) (string, error) {
	p := string(password[:])
	if username != "Puspa" || p != "1234" {
		return "", errors.New("Username/Password is wrong!")
	}

	return "Succesfully logged in!", nil
}
