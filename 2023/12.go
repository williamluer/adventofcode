package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func springFit(springs []string, groups []int) bool {

	groupsIter := 0
	currentSpringGroupSize := 0
	for i := 0; i < len(springs); i++ {
		// fmt.Println(springs, groups, groupsIter, currentSpringGroupSize, i, springs[:i])
		if springs[i] == "?" {
			// fmt.Println("here")
			return false
		} else if springs[i] == "#" {
			// fmt.Println("its a #")
			currentSpringGroupSize++
			if i == len(springs)-1 {
				if groupsIter == len(groups)-1 && groups[groupsIter] == currentSpringGroupSize {
					// fmt.Println("mmmm")
					// fmt.Println(springs)
					return true
				} else {
					// fmt.Println("pppp")
					return false
				}
			}
		} else if springs[i] == "." {
			// fmt.Println("its a .")
			if currentSpringGroupSize == 0 {
			} else if groupsIter < len(groups) && groups[groupsIter] == currentSpringGroupSize {
				currentSpringGroupSize = 0
				groupsIter++
			} else {
				return false
			}
		} else {
			fmt.Println("\n\n\nSOMETHING UNEXPECTED HAPPENED\n\n\n")
			return false
		}
	}
	// fmt.Println(springs)

	// fmt.Println("SF:SDFJDSLFJ")
	return groupsIter >= len(groups)
}

func findMatch(springs []string, groups []int, i int) int {
	// fmt.Println("Checking ", springs, groups, i)

	firstPath := 0
	secondPath := 0
	if slices.Contains(springs, "?") {
		if springs[i] == "?" {
			newSprings1 := make([]string, len(springs))
			copy(newSprings1, springs)
			newSprings1[i] = "."
			firstPath = findMatch(newSprings1, groups, i+1)
			newSprings2 := make([]string, len(springs))
			copy(newSprings2, springs)
			newSprings2[i] = "#"
			secondPath = findMatch(newSprings2, groups, i+1)
			return firstPath + secondPath

		}
		return findMatch(springs, groups, i+1)
	} else {
		if springFit(springs, groups) {
			fmt.Println("Returning 1 for ", springs, groups)
			return 1
		}
		return 0
	}
}

func main() {
	filename := "12-sample-2.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	groups := make([][]int, 0)
	springs := make([][]string, 0)
	for scanner.Scan() {
		springRe := regexp.MustCompile(`[.?#]`)
		groupsRe := regexp.MustCompile(`\d+`)
		groupsSt := make([]string, 0)
		groupsInt := make([]int, 0)
		current_line := scanner.Text()
		springSpecs := springRe.FindAllString(current_line, -1)
		springs = append(springs, springSpecs)
		groupsSt = groupsRe.FindAllString(current_line, -1)
		for i := 0; i < len(groupsSt); i++ {
			val, _ := strconv.Atoi(groupsSt[i])
			groupsInt = append(groupsInt, val)
		}
		groups = append(groups, groupsInt)
	}
	fmt.Println(groups)
	fmt.Println(springs)

	count := 0
	for i := 0; i < len(springs); i++ {
		c := findMatch(springs[i], groups[i], 0)
		fmt.Println("Count for ", springs[i], groups[i], c)
		count += c
	}
	fmt.Println(count)

	// for i := 0; i < len(springs); i++ {
	// 	fmt.Println("CHECKING and ", springFit(springs[i], groups[i]))
	// }
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
