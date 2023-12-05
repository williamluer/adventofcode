package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	"github.com/juliangruber/go-intersect"
)

func main() {
	filename := "4.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	sum := 0
	i := 0
	for scanner.Scan() {
		currentLine := scanner.Text()
		numbers := re.FindAllString(currentLine, -1)
		winningNumbers := numbers[1:11]
		guessedNumbers := numbers[11:]
		fmt.Println(winningNumbers, guessedNumbers)
		intersection := intersect.Simple(winningNumbers, guessedNumbers)
		nWins := len(intersection)
		fmt.Println(intersection)
		fmt.Println(nWins)
		if nWins > 0 {
			points := math.Pow(2, float64(nWins-1))
			fmt.Println(points)
			sum += int(points)
		}
		fmt.Println()

	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
