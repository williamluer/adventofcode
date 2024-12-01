package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// open the file using Open() function from os library
	filename := "1.txt"
	sum := 0
	re := `\d`
	r := regexp.MustCompile(re)

	int_strs := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var current_line string = ""

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		current_line = scanner.Text()

		fmt.Println(current_line)
		for i := 0; i < len(int_strs); i++ {
			current_line = strings.Replace(current_line, int_strs[i], int_strs[i]+strconv.Itoa(i+1)+int_strs[i], -1)
		}
		fmt.Println(current_line)
		matches := r.FindAllString(current_line, -1)
		full_int := matches[0] + matches[len(matches)-1]
		val1, _ := strconv.Atoi(full_int)
		fmt.Println("Got ", val1)
		sum += val1
		fmt.Println(sum)
		fmt.Println()
		// fmt.Println(val1)
		// fmt.Println(matches)
		// fmt.Println(sum)
		// fmt.Println()

		// current_line.
		// Replace(current_line,
		// replacements =
	}
	// check for the error that occurred during the scanning

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the file at the end of the program
	file.Close()

	fmt.Println(sum)
}
