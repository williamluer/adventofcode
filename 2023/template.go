package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	filename := "2.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reg := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current_line := scanner.Text()
		fmt.Println(current_line)
		// do something
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
