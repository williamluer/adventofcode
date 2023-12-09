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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func isFinished(nodes []string) bool {
	for _, n := range nodes {
		if !strings.HasSuffix(n, "Z") {
			return false
		}
	}
	fmt.Println("Returning true for ", nodes)
	return true
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
	count := 0

	// for !isFinished(currentNodes) {
	// 	count += 1
	// 	for j, currentNode := range currentNodes {
	// 		currentTurn := directions[(count-1)%len(directions)]
	// 		if currentTurn == "L" {
	// 			currentNodes[j] = nodeMap[currentNode].left
	// 		}
	// 		if currentTurn == "R" {
	// 			currentNodes[j] = nodeMap[currentNode].right
	// 		}
	// 	}
	// }

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
