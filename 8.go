package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func main() {
	filename := "8.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reg := regexp.MustCompile(`[A-Z]{3}`)
	scanner := bufio.NewScanner(file)

	var directions []string
	nodeMap := make(map[string]Node)

	i := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		fmt.Println(current_line)
		if i == 0 {
			for _, r := range current_line {
				directions = append(directions, string(r))
			}
			fmt.Println(directions)
		}
		if i > 1 {
			nodes := reg.FindAllString(current_line, -1)
			node := Node{left: nodes[1], right: nodes[2]}
			nodeMap[nodes[0]] = node
		}
		i++
	}
	fmt.Println(nodeMap)
	currentNode := "AAA"
	endingNode := "ZZZ"

	i = 0
	count := 0
	for currentNode != endingNode {
		fmt.Println("Current node ", currentNode)
		count += 1
		currentTurn := directions[i%len(directions)]
		fmt.Println(currentTurn)
		if currentTurn == "L" {
			currentNode = nodeMap[currentNode].left
		}
		if currentTurn == "R" {
			currentNode = nodeMap[currentNode].right
		}
		i++
	}
	fmt.Println("Steps ", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
