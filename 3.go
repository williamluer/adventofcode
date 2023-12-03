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
		for currentCol, currentVal := range matches[currentRow] { // matchIndex is [[0 1] [3 6]]
			startIndex := matchIndices[currentRow][currentCol][0]
			endIndex := matchIndices[currentRow][currentCol][1]

			useIt := false
		out:
			for rowOffset := -1; rowOffset <= 1; rowOffset++ {
				for colOffset := startIndex - 1; colOffset <= endIndex; colOffset++ {
					if containsSpecialChars(spec[currentRow+rowOffset][colOffset]) {
						fmt.Println(currentVal, " matched ", spec[currentRow+rowOffset][colOffset])
						useIt = true
						break out
					}
				}
			}
			if useIt {
				fmt.Println("adding ", getInt(currentVal), " to ", sum)
				sum += getInt(currentVal)
			}
			if !useIt {
				fmt.Println(currentVal, " didn't match")
			}
		}
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

}
