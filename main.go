package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validateAge(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	if num < 0 {
		return 0, errors.New("negative")
	}

	return num, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	num, err := validateAge(s)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", num)
	}

}
