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

func firstCharacter(s string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(s)
	i, _ := strconv.Atoi(match)
	return i
}

func main() {
	filename := "2.txt"
	max_r := 12
	max_g := 13
	max_b := 14
	sum := 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	red_r := regexp.MustCompile(`\d+ red`)
	blue_r := regexp.MustCompile(`\d+ blue`)
	green_r := regexp.MustCompile(`\d+ green`)

	i := 0
	scanner := bufio.NewScanner(file)
	running_power := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		valid := true
		fmt.Println(current_line)
		game := strings.Split(current_line, ":")
		rounds := strings.Split(game[1], ";")

		round_max_r := 0
		round_max_g := 0
		round_max_b := 0
		for i := 0; i < len(rounds); i++ {
			round := rounds[i]
			current_red := firstCharacter(red_r.FindString(round))
			current_blue := firstCharacter(blue_r.FindString(round))
			current_green := firstCharacter(green_r.FindString(round))
			if current_red > round_max_r {
				round_max_r = current_red
			}
			if current_green > round_max_g {
				round_max_g = current_green
			}
			if current_blue > round_max_b {
				round_max_b = current_blue
			}
			fmt.Println("red", current_red)
			fmt.Println("green", current_green)
			fmt.Println("blue", current_blue)
			if current_red > max_r || current_blue > max_b || current_green > max_g {
				valid = false
				fmt.Println("Not possible")
			}
		}
		running_power += round_max_r * round_max_g * round_max_b
		fmt.Println("Doing power with", round_max_r, round_max_g, round_max_b)
		if valid {
			fmt.Println("Possible. adding", i+1, " to ", sum)
			sum += i + 1
		}
		fmt.Println("")

		i += 1
	}
	fmt.Println("Sum of game IDs", sum)
	fmt.Println("Power", running_power)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
