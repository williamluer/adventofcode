package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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
	return groupsIter >= len(groups)
}

var memo map[string]int

func isValidPath(springs []string, groups []int, groupSum int) bool {
	sumQsAndSprings := 0
	groupIndex := 0
	currentGroupSizeCounter := 0

	for _, s := range springs {
		switch s {
		case "#":
			currentGroupSizeCounter++
			if currentGroupSizeCounter > groups[groupIndex] {
				return false
			}
			sumQsAndSprings++

		case ".":
			if currentGroupSizeCounter != 0 {
				if currentGroupSizeCounter < groups[groupIndex] {
					return false
				}
				currentGroupSizeCounter = 0
				groupIndex++
				if groupIndex >= len(groups) {
					return sumQsAndSprings >= groupSum
				}
			}

		case "?":
			return true
		}
	}

	return sumQsAndSprings >= groupSum && groupIndex >= len(groups)-1
}

// func findMatch(springs []string, groups []int, groupsSum int, i int) int {
// 	// fmt.Println("Checking ", springs, groups, i)

// 	springsStr := strings.Join(springs, "")
// 	key := MemoKey{Springs: springsStr, Index: i}

// 	// Check if result is in the cache
// 	if val, found := memo[key]; found {
// 		return val
// 	}

// 	firstPath := 0
// 	secondPath := 0
// 	if slices.Contains(springs, "?") {
// 		if springs[i] == "?" {
// 			newSprings1 := make([]string, len(springs))
// 			copy(newSprings1, springs)
// 			newSprings1[i] = "."
// 			if isValidPath(newSprings1, groups, groupsSum) {
// 				firstPath = findMatch(newSprings1, groups, groupsSum, i+1)
// 			}
// 			newSprings2 := make([]string, len(springs))
// 			copy(newSprings2, springs)
// 			newSprings2[i] = "#"

//				if isValidPath(newSprings2, groups, groupsSum) {
//					secondPath = findMatch(newSprings2, groups, groupsSum, i+1)
//				}
//				result := firstPath + secondPath
//				memo[springsStr] = result
//				return result
//			}
//			result := findMatch(springs, groups, groupsSum, i+1)
//			memo[key] = result
//			return result
//		} else {
//			if springFit(springs, groups) {
//				// fmt.Println("Returning 1 for ", springs, groups)
//				return 1
//			}
//			return 0
//		}
//	}
func findMatch(springs []string, groups []int, groupsSum int, i int, memo map[string]int) int {
	fmt.Println(memo)
	// springsStr := strings.Join(springs, "")
	// key := MemoKey{Springs: springsStr, Index: i}

	// // Check if result is in the cache
	// if val, found := memoKeys[key]; found {
	// 	return val
	// }
	key := fmt.Sprintf("%v-%v-%d", strings.Join(springs, ""), strings.Join(convertToStringArray(groups), "-"), i)
	// Check if the result is already in the cache
	if val, ok := memo[key]; ok {
		return val
	}

	firstPath, secondPath := 0, 0
	if slices.Contains(springs, "?") {
		if springs[i] == "?" {
			newSprings1 := make([]string, len(springs))
			copy(newSprings1, springs)
			newSprings1[i] = "."
			if isValidPath(newSprings1, groups, groupsSum) {
				firstPath = findMatch(newSprings1, groups, groupsSum, i+1, memo)
			}

			newSprings2 := make([]string, len(springs))
			copy(newSprings2, springs)
			newSprings2[i] = "#"
			if isValidPath(newSprings2, groups, groupsSum) {
				secondPath = findMatch(newSprings2, groups, groupsSum, i+1, memo)
			}
		} else {
			return findMatch(springs, groups, groupsSum, i+1, memo)
		}
	} else {
		if springFit(springs, groups) {
			// memoKeys[key] = 1
			return 1
		}
		// memoKeys[key] = 0
		return 0
	}

	result := firstPath + secondPath
	// memoKeys[key] = result
	memo[key] = result

	return result
}
func convertToStringArray(ints []int) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strs
}

func main() {
	filename := "12-sample.txt"
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
		ogSpecs := make([]string, len(springSpecs))

		copy(ogSpecs, springSpecs)

		for i := 0; i < 4; i++ {
			springSpecs = append(springSpecs, "?")
			springSpecs = append(springSpecs, ogSpecs...)
		}
		springs = append(springs, springSpecs)
		groupsSt = groupsRe.FindAllString(current_line, -1)
		for i := 0; i < len(groupsSt); i++ {
			val, _ := strconv.Atoi(groupsSt[i])
			groupsInt = append(groupsInt, val)
		}
		ogGroups := make([]int, len(groupsInt))

		copy(ogGroups, groupsInt)
		for i := 0; i < 4; i++ {
			groupsInt = append(groupsInt, ogGroups...)
		}

		groups = append(groups, groupsInt)
	}
	// fmt.Println(groups)
	// fmt.Println(springs)

	count := 0
	counts := make([]int, 0)
	for i := 0; i < len(springs); i++ {
		groupCount := 0
		for g := 0; g < len(groups[i]); g++ {
			groupCount += groups[i][g]
		}
		// fmt.Println("Group count", groupCount)
		memo = make(map[string]int)

		c := findMatch(springs[i], groups[i], groupCount, 0, memo)
		fmt.Println("Count for ", springs[i], groups[i], c)
		count += c
		counts = append(counts, c)
		fmt.Println(counts)
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
