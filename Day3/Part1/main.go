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

	hitAmt := CheckTreeCollisions(0, geoMap)

	fmt.Printf("Hit %d trees\n", hitAmt)

	elapsed := time.Since(start)
	fmt.Printf("Execution Time: %s\n", elapsed)
	os.Exit(1)
}

func CheckTreeCollisions(startIndex int, geoMap [][]string) int {
	rightLoc := 0
	downAmt := 1
	rightAmt := 3
	hitAmt := 0
	for i := startIndex; i < len(geoMap); i = i + downAmt {
		if geoMap[i][rightLoc] == "#" {
			hitAmt++
		}
		rightLoc = rightLoc + rightAmt
	}
	return hitAmt
}
