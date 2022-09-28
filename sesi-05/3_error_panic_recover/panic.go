package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

	fmt.Println(strings.Repeat("-", 25))
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}
