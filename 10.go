package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var pipeGridTmp [][]string
var pipeGrid [][]string

func findNextDirection(grid [][]string, r int, c int, visited map[int]bool) (int, int) {
	if grid[r][c-1] == "-" && (grid[r][c] == "7" || grid[r][c] == "J" || grid[r][c] == "-" || grid[r][c] == "S") { // WEST
		// fmt.Println("Next path: -")
		return r, c - 1
	} else if grid[r][c-1] == "L" && (grid[r][c] == "7" || grid[r][c] == "J" || grid[r][c] == "-" || grid[r][c] == "S") {
		// fmt.Println("Next path: L")
		return r, c - 1
	} else if grid[r][c-1] == "F" && (grid[r][c] == "7" || grid[r][c] == "J" || grid[r][c] == "-" || grid[r][c] == "S") {
		// fmt.Println("Next path: F")
		return r, c - 1
	} else if grid[r][c-1] == "X" && len(visited) > 2 {
		// fmt.Println("Next path: S")
		return r, c - 1
	} else if grid[r][c+1] == "-" && (grid[r][c] == "-" || grid[r][c] == "L" || grid[r][c] == "F" || grid[r][c] == "S") { // EAST
		// fmt.Println("Next path: -")
		return r, c + 1
	} else if grid[r][c+1] == "J" && (grid[r][c] == "-" || grid[r][c] == "L" || grid[r][c] == "F" || grid[r][c] == "S") {
		// fmt.Println("Next path: J")
		return r, c + 1
	} else if grid[r][c+1] == "7" && (grid[r][c] == "-" || grid[r][c] == "L" || grid[r][c] == "F" || grid[r][c] == "S") {
		// fmt.Println("Next path: 7")
		return r, c + 1
	} else if grid[r][c+1] == "X" && len(visited) > 2 {
		// fmt.Println("Next path: S")
		return r, c + 1
	} else if grid[r+1][c] == "|" && (grid[r][c] == "|" || grid[r][c] == "7" || grid[r][c] == "F" || grid[r][c] == "S") { // SOUTH
		// fmt.Println("Next path: |")
		return r + 1, c
	} else if grid[r+1][c] == "L" && (grid[r][c] == "|" || grid[r][c] == "7" || grid[r][c] == "F" || grid[r][c] == "S") {
		// fmt.Println("Next path: L")
		return r + 1, c
	} else if grid[r+1][c] == "J" && (grid[r][c] == "|" || grid[r][c] == "7" || grid[r][c] == "F" || grid[r][c] == "S") {
		// fmt.Println("Next path: J")
		return r + 1, c
	} else if grid[r+1][c] == "X" && len(visited) > 2 {
		// fmt.Println("Next path: S")
		return r + 1, c
	} else if grid[r-1][c] == "|" && (grid[r][c] == "|" || grid[r][c] == "L" || grid[r][c] == "J" || grid[r][c] == "S") { // NORTH
		// fmt.Println("Next path: |")
		return r - 1, c
	} else if grid[r-1][c] == "7" && (grid[r][c] == "|" || grid[r][c] == "L" || grid[r][c] == "J" || grid[r][c] == "S") {
		// fmt.Println("Next path: 7")
		return r - 1, c
	} else if grid[r-1][c] == "F" && (grid[r][c] == "|" || grid[r][c] == "L" || grid[r][c] == "J" || grid[r][c] == "S") {
		// fmt.Println("Next path: F")
		return r - 1, c
	} else if grid[r-1][c] == "X" && len(visited) > 2 {
		// fmt.Println("Next path: S")
		return r - 1, c
	}
	return -10000000000, -10000000000
}

func findMaxPath(grid [][]string, r int, c int, visited map[int]bool, visitedWalls map[int]bool, visitedL map[int]bool, visited7 map[int]bool, visitedJ map[int]bool, visitedF map[int]bool) int {
	visited[r*len(grid[0])+c] = true
	if grid[r][c] == "|" {
		visitedWalls[r*len(grid[0])+c] = true
	}
	if grid[r][c] == "L" {
		visitedL[r*len(grid[0])+c] = true

	}
	if grid[r][c] == "7" {
		visited7[r*len(grid[0])+c] = true
	}
	if grid[r][c] == "J" {
		visitedJ[r*len(grid[0])+c] = true
	}
	if grid[r][c] == "F" {
		visitedF[r*len(grid[0])+c] = true
	}

	newRow, newCol := findNextDirection(grid, r, c, visited)

	// for r := 1; r < len(grid)-1; r++ {
	// 	for c := 1; c < len(grid[0])-1; c++ {
	// 		fmt.Print(grid[r][c])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	if len(visited) == 1 {
		grid[r][c] = "X"
	} else {
		grid[r][c] = "0"
	}
	_, ok := visited[r*len(grid[0])+c]
	if grid[newRow][newCol] == "X" || !ok {
		return 1
	} else {
		return 1 + findMaxPath(grid, newRow, newCol, visited, visitedWalls, visitedL, visited7, visitedJ, visitedF)
	}
}

func isWall(s string) bool {
	return s == "┃" || s == "┗" || s == "┛" //|| s == "┏" || s == "┓"
}

func main() {
	filename := "10.txt"
	visited := make(map[int]bool)
	visitedWalls := make(map[int]bool)
	visitedL := make(map[int]bool)
	visited7 := make(map[int]bool)
	visitedJ := make(map[int]bool)
	visitedF := make(map[int]bool)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reg := regexp.MustCompile(`[\|\-FJ7L.S]`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current_line := scanner.Text()
		pipeLine := reg.FindAllString(current_line, -1)
		pipeGridTmp = append(pipeGridTmp, pipeLine)
	}
	l := len(pipeGridTmp[0])

	for i := 0; i <= len(pipeGridTmp)+1; i++ {
		if i == 0 || i == len(pipeGridTmp)+1 {
			buffer := make([]string, l+2)
			for j, _ := range buffer {
				buffer[j] = "."
			}
			pipeGrid = append(pipeGrid, buffer)
		} else {
			paddedSlice := []string{"."}
			paddedSlice = append(paddedSlice, pipeGridTmp[i-1]...)
			paddedSlice = append(paddedSlice, ".")
			pipeGrid = append(pipeGrid, paddedSlice)
		}
	}

	maxPath := 0
	for r := 1; r < len(pipeGrid)-1; r++ {
		for c := 1; c < len(pipeGrid[0])-1; c++ {
			if pipeGrid[r][c] == "S" {
				maxPath = findMaxPath(pipeGrid, r, c, visited, visitedWalls, visitedL, visited7, visitedJ, visitedF)
			}
		}
	}
	fmt.Println(maxPath)

	const gridSize = 142
	var visitedGrid [gridSize][gridSize]string
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			visitedGrid[r][c] = "."
		}
	}

	for k, _ := range visited {
		row := k / gridSize
		col := k % gridSize
		visitedGrid[row][col] = "━"
	}
	for k, _ := range visitedWalls {
		row := k / gridSize
		col := k % gridSize
		visitedGrid[row][col] = "┃"
	}
	for k, _ := range visitedL {
		row := k / gridSize
		col := k % gridSize
		// Corner character
		visitedGrid[row][col] = "┗"
	}
	for k, _ := range visited7 {
		row := k / gridSize
		col := k % gridSize
		// Corner character
		visitedGrid[row][col] = "┓"
	}
	for k, _ := range visitedJ {
		row := k / gridSize
		col := k % gridSize
		visitedGrid[row][col] = "┛"
	}
	for k, _ := range visitedF {
		row := k / gridSize
		col := k % gridSize
		visitedGrid[row][col] = "┏"
	}
	// fmt.Println(visitedGrid)
	enclosedSpaces := 0
	type Spot struct {
		row int
		col int
	}

	for r := 1; r < len(visitedGrid)-1; r++ {
		nWalls := 0
		tentativeSpaces := 0
		tentativeSpaceList := make([]Spot, 0)
		for c := 1; c < len(visitedGrid[0])-1; c++ {
			if isWall(visitedGrid[r][c]) {
				nWalls++
			}
			if nWalls%2 == 1 && visitedGrid[r][c] == "." {
				tentativeSpaces++
				tentativeSpaceList = append(tentativeSpaceList, Spot{row: r, col: c})
			}
			if nWalls%2 == 0 {
				enclosedSpaces += tentativeSpaces
				tentativeSpaces = 0
				for _, s := range tentativeSpaceList {
					visitedGrid[s.row][s.col] = "I"
				}
				tentativeSpaceList = []Spot{}
			}
		}
		tentativeSpaceList = []Spot{}
	}
	for r := 1; r < len(visitedGrid)-1; r++ {
		for c := 1; c < len(visitedGrid[0])-1; c++ {
			fmt.Print(visitedGrid[r][c])
		}
		fmt.Println()
	}
	fmt.Println("Enclosed spaces ", enclosedSpaces)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()

}
