package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	numberOfValid := 0
	dataString := string(data)
	s := strings.Split(dataString, "\n")

	for _, passwordInput := range s {
		passwordMap := strings.Split(passwordInput, " ")
		indexes := strings.Split(passwordMap[0], "-")
		min, _ := strconv.Atoi(indexes[0])
		max, _ := strconv.Atoi(indexes[1])
		char := passwordMap[1][0]
		if isValidPassword(passwordMap[2], min, max, int(char)) {
			numberOfValid = numberOfValid + 1
		}
	}

	fmt.Printf("%d passwords are valid.\n", numberOfValid)

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
	os.Exit(1)
}

func isValidPassword(password string, minAmt int, maxAmt int, characterToCheck int) bool {
	countOfChars := 0
	for _, char := range password {
		if characterToCheck == int(char) {
			countOfChars = countOfChars + 1
		}
	}
	if countOfChars >= minAmt && countOfChars <= maxAmt {
		return true
	}

	return false
}
