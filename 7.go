package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
)

type Hand struct {
	cards    []int
	bid      int
	bestHand int // 6 for five of a kind, 5 for four of a kind,... 0 for high card
}

const FIVE_KIND = 6
const FOUR_KIND = 5
const FULL_HOUSE = 4
const THREE_KIND = 3
const TWO_PAIR = 2
const ONE_PAIR = 1
const HIGH_CARD = 0

func rankHands(h1 Hand, h2 Hand) bool {
	// returns true if h1 is a worse hand than h2
	if h1.bestHand > h2.bestHand {
		return true
	} else if h1.bestHand < h2.bestHand {
		return false
	} else {
		for i := 0; i < 5; i++ {
			if h1.cards[i] > h2.cards[i] {
				return true
			} else if h1.cards[i] < h2.cards[i] {
				return false
			}
		}
	}
	return false
}

func checkForTwoPair(cardCount []int) bool {
	nPairs := 0
	for _, c := range cardCount {
		if c == 2 {
			nPairs++
		}
	}
	return nPairs == 2
}

func findBestHand(cards []int) int {
	cardMap := make([]int, 15)
	for _, c := range cards {
		cardMap[c]++
	}
	if slices.Contains(cardMap, 5) {
		return FIVE_KIND
	} else if slices.Contains(cardMap, 4) {
		return FOUR_KIND
	} else if slices.Contains(cardMap, 3) && slices.Contains(cardMap, 2) {
		return FULL_HOUSE
	} else if slices.Contains(cardMap, 3) {
		return THREE_KIND
	} else if checkForTwoPair(cardMap) {
		return TWO_PAIR
	} else if slices.Contains(cardMap, 2) {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}

}

func main() {
	filename := "7.txt"
	cardToIntMap := map[string]int{
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	handRegex := regexp.MustCompile(`[^\s]+`)
	cardsRegex := regexp.MustCompile(`.`)

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current_line := scanner.Text()
		cardsBids := handRegex.FindAllString(current_line, -1)
		cards := cardsRegex.FindAllString(cardsBids[0], -1)
		bidInt, _ := strconv.Atoi(cardsBids[1])
		var cardInts []int
		for _, c := range cards {
			cardInts = append(cardInts, cardToIntMap[c])
		}
		hand := Hand{cards: cardInts, bid: bidInt, bestHand: findBestHand(cardInts)}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return !rankHands(hands[i], hands[j])
	})
	answer := 0
	for i, hand := range hands {
		answer += (i + 1) * hand.bid
	}
	fmt.Println("Answer: ", answer)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
