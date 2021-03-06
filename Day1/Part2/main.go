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

	gotAnswer := false
	dataString := string(data)
	s := strings.Split(dataString, "\n")

	for _, strNum1 := range s {
		num1, _ := strconv.Atoi(strNum1)

		for _, strNum2 := range s {
			num2, _ := strconv.Atoi(strNum2)

			for _, strNum3 := range s {

				num3, _ := strconv.Atoi(strNum3)

				if num1+num2+num3 == 2020 {
					gotAnswer = true

					fmt.Printf("Number 1 is %d, Number 2 is %d, Number 3 is %d\n", num1, num2, num3)

					answer := num1 * num2 * num3

					fmt.Printf("Answer is %d\n", answer)

					break
				}
			}

			if gotAnswer {
				break
			}
		}

		if gotAnswer {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution Time: %s", elapsed)
	os.Exit(1)
}
