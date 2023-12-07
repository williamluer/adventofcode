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
	i := 1
	for scanner.Scan() {
		currentLine := scanner.Text()
		numbers := re.FindAllString(currentLine, -1)
		winningNumbers := numbers[1:11]
		guessedNumbers := numbers[11:]
		intersection := intersect.Simple(winningNumbers, guessedNumbers)
		nWins := len(intersection)
		fmt.Println(i, " has ", nWins)
		if nWins > 0 {
			points := math.Pow(2, float64(nWins-1))
			sum += int(points)
		}
		i += 1
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
