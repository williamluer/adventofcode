package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func checkInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
func getInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func containsSpecialChars(s string) bool {
	pattern := "[^a-zA-Z0-9.]"
	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	// fmt.Println(s, r.MatchString(s))
	return r.MatchString(s)
}

func main() {
	gearMap := make(map[string][]int)
	filename := "3.txt"
	file, err := os.Open(filename)
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
	}

	re := `\d+`
	reg := regexp.MustCompile(re)

	var matches [142][]string
	var matchIndices [142][][]int

	spec := [142][142]string{}
	r := 1
	for scanner.Scan() {
		currentLine := scanner.Text()
		x := reg.FindAllString(currentLine, -1)
		y := reg.FindAllStringIndex(currentLine, -1)
		for i, v := range y {
			for j, w := range v {
				y[i][j] = w + 1
			}
		}
		matches[r] = x
		matchIndices[r] = y
		for c, ch := range currentLine {
			spec[r][c+1] = string(ch)
		}
		r += 1
	}
	sum := 0
	for currentRow := 1; currentRow < 141; currentRow++ {
		fmt.Println(gearMap)
		for currentCol, currentVal := range matches[currentRow] { // matchIndex is [[0 1] [3 6]]
			startIndex := matchIndices[currentRow][currentCol][0]
			endIndex := matchIndices[currentRow][currentCol][1]
			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				for colOffset := startIndex - 1; colOffset <= endIndex; colOffset++ {
					if spec[currentRow+rowOffset][colOffset] == "*" {
						key := strconv.Itoa(currentRow+rowOffset) + "-" + strconv.Itoa(colOffset)
						gearMap[key] = append(gearMap[key], getInt(currentVal))
						// fmt.Println(currentVal, " matched ", spec[currentRow+rowOffset][colOffset])
					}
				}
			}
		}
	}

	sum = 0
	for indices, vals := range gearMap {
		runningMultiplier := 1
		if len(vals) >= 2 {
			for i := 0; i < len(vals); i++ {
				runningMultiplier *= vals[i]
			}
			sum += runningMultiplier
		}
		fmt.Println(indices, " ", vals, " ", runningMultiplier)
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

}
