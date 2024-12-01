package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time     int
	distance int
}

func calculateDistance(holdTime int, totalTime int) int {
	// fmt.Println("Distance for ", holdTime, " and totalTime ", totalTime, "is ", holdTime*(totalTime-holdTime))
	return holdTime * (totalTime - holdTime)
}

func main() {
	filename := "6-2.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	nRaces := 1
	races := make([]Race, nRaces)
	reg := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		parsedNumbers := reg.FindAllString(current_line, -1)
		fmt.Println(parsedNumbers)
		for j := 0; j < len(parsedNumbers); j++ {
			if i == 0 {
				val, _ := strconv.Atoi(parsedNumbers[j])
				races[j].time = val
			} else if i == 1 {
				val, _ := strconv.Atoi(parsedNumbers[j])
				races[j].distance = val
			}
		}
		i += 1
		fmt.Println(races)

	}
	fmt.Println(races)
	scores := make([]int, nRaces)
	for j, race := range races {
		fmt.Println("ON race ", race)
		waysToWin := 0
		for i := 0; i < race.time; i++ {
			if calculateDistance(i, race.time) > race.distance {
				waysToWin++
			}
		}
		scores[j] = waysToWin
		fmt.Println(scores)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
