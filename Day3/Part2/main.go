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

	//Read the map.
	dataString := string(data)
	xLines := strings.Split(dataString, "\n")

	geoMap := [][]string{}

	for _, xLine := range xLines {
		yChars := []string{}
		outYChars := []string{}
		for _, yChar := range xLine {
			yChars = append(yChars, string(yChar))
		}
		duplicationAmount := 200
		for i := 0; i < duplicationAmount; i++ {
			outYChars = append(outYChars, yChars...)
		}

		geoMap = append(geoMap, outYChars)
	}

	//Check the required deltas
	deltas := [][]int{}
	deltas = append(deltas, []int{1, 1})
	deltas = append(deltas, []int{1, 3})
	deltas = append(deltas, []int{1, 5})
	deltas = append(deltas, []int{1, 7})
	deltas = append(deltas, []int{2, 1})

	totalAmt := CheckTreeCollisions(0, deltas[0][0], deltas[0][1], geoMap)
	for i := 1; i < len(deltas); i++ {
		totalAmt = totalAmt * CheckTreeCollisions(0, deltas[i][0], deltas[i][1], geoMap)
	}

	fmt.Printf("Answer is %d trees\n", totalAmt)

	elapsed := time.Since(start)
	fmt.Printf("Execution Time: %s\n", elapsed)
	os.Exit(1)
}

func CheckTreeCollisions(startIndex int, downAmt int, rightAmt int, geoMap [][]string) int {
	rightLoc := 0
	hitAmt := 0
	for i := startIndex; i < len(geoMap); i = i + downAmt {
		if geoMap[i][rightLoc] == "#" {
			hitAmt++
		}
		rightLoc = rightLoc + rightAmt
	}
	return hitAmt
}
