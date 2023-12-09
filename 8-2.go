package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
	reg := regexp.MustCompile(`[0-9A-Z]{3}`)
	scanner := bufio.NewScanner(file)

	var directions []string
	var startingNodes []string
	nodeMap := make(map[string]Node)

	i := 0
	for scanner.Scan() {
		current_line := scanner.Text()
		if i == 0 {
			for _, r := range current_line {
				directions = append(directions, string(r))
			}
		}
		if i > 1 {
			nodes := reg.FindAllString(current_line, -1)
			currNode := Node{left: nodes[1], right: nodes[2]}
			nodeMap[nodes[0]] = currNode
			if strings.HasSuffix(nodes[0], "A") {
				startingNodes = append(startingNodes, nodes[0])
			}
		}
		i++
	}
	currentNodes := startingNodes

	var counts []int
	for _, currentNode := range currentNodes {
		count := 0
		for !strings.HasSuffix(currentNode, "Z") {
			count += 1
			currentTurn := directions[(count-1)%len(directions)]
			if currentTurn == "L" {
				currentNode = nodeMap[currentNode].left
			}
			if currentTurn == "R" {
				currentNode = nodeMap[currentNode].right
			}
			i++
		}
		counts = append(counts, count)
	}
	fmt.Println("counts ", counts)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
