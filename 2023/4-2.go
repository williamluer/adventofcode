package main

import (
	"bufio"
	"fmt"
	"log"
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
	cardCounter := make([]int, 220)
	for i := range cardCounter {
		cardCounter[i] = 1
	}

	for scanner.Scan() {
		currentLine := scanner.Text()
		fmt.Println(cardCounter)
		for cardCount := cardCounter[i]; cardCount >= 1; cardCount-- {
			numbers := re.FindAllString(currentLine, -1)
			winningNumbers := numbers[1:11]
			guessedNumbers := numbers[11:]
			intersection := intersect.Simple(winningNumbers, guessedNumbers)
			nWins := len(intersection)
			for incCard := nWins; incCard > 0; incCard-- {
				if i+nWins < 219 {
					cardCounter[i+incCard] += 1
				}
			}
		}
		i += 1
	}
	for i := range cardCounter {
		sum += cardCounter[i]
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
