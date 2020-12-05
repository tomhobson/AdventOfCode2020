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
		index1, _ := strconv.Atoi(indexes[0])
		index1--
		index2, _ := strconv.Atoi(indexes[1])
		index2--
		char := passwordMap[1][0]
		if isValidPassword(passwordMap[2], index1, index2, char) {
			numberOfValid = numberOfValid + 1
		}
	}

	fmt.Printf("%d passwords are valid.\n", numberOfValid)

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
	os.Exit(1)
}

func isValidPassword(password string, index1 int, index2 int, characterToCheck byte) bool {
	countOfChars := 0
	if password[index1] == characterToCheck {
		countOfChars++
	}

	if password[index2] == characterToCheck {
		countOfChars++
	}

	if countOfChars == 1 {
		return true
	}

	return false
}
