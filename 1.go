package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "regexp"
    "strconv"
)

func main() {
	// open the file using Open() function from os library
	filename := "1.txt"
    sum := 0 
    re := `\d`
    r := regexp.MustCompile(re)

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
        matches := r.FindAllString(current_line,-1)
        full_int := matches[0] + matches[len(matches)-1]
        val1, _ := strconv.Atoi(full_int)
        sum += val1 
        fmt.Println(val1)
        fmt.Println(matches)
        fmt.Println(sum)
        fmt.Println()
	}
    // check for the error that occurred during the scanning
    
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    
	// Close the file at the end of the program
	file.Close()

    fmt.Println(sum)
}


