package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	dataString := string(data)
	s := strings.Split(dataString, "\n")

	for _, passwordInput := range s {
		passwordMap := strings.Split(passwordInput, " ")
		minMax := strings.Split(passwordMap[0], "-")
		min := minMax[0]
		max := minMax[1]
		char := passwordMap[1][0]
		fmt.Printf("Min Length: %s Max Length: %s, Character to check: %s, Password: %s\n", min, max, string(char), passwordMap[2])
	}

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
	os.Exit(1)
}
