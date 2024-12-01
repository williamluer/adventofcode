package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
)

func printGrid(grid [][]string) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			fmt.Print(grid[r][c])
		}
		fmt.Println()
	}
}

func manhattanDistance(p1 Point, p2 Point) int {
	return int(math.Abs(float64(p1.x-p2.x))) + int(math.Abs(float64(p1.y-p2.y)))
}
func duplicateRows(grid [][]string) [][]string {
	for r := 0; r < len(grid); r++ {
		if slices.Index(grid[r], "#") == -1 {
			fmt.Println("Inserting at row", r)
			grid = slices.Insert(grid, r, grid[r])
			r++
		}
	}
	return grid
}

func duplicateCols(grid [][]string) [][]string {
	for c := 0; c < len(grid[0]); c++ {
		duplicateCol := true
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == "#" {
				duplicateCol = false
			}
		}
		if duplicateCol {
			c++
			for r := 0; r < len(grid); r++ {
				grid[r] = slices.Insert(grid[r], c, ".")
			}
		}
	}
	return grid
}

type Point struct {
	x int
	y int
}

func main() {
	filename := "11.txt"
	file, err := os.Open(filename)
	galaxies := make([]Point, 0)
	if err != nil {
		log.Fatal(err)
	}
	universeGrid := make([][]string, 0)
	reg := regexp.MustCompile(`[\.#]`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current_line := scanner.Text()
		row := reg.FindAllString(current_line, -1)
		universeGrid = append(universeGrid, row)
	}
	universeGrid = duplicateRows(universeGrid)
	universeGrid = duplicateCols(universeGrid)
	printGrid(universeGrid)

	for r := 0; r < len(universeGrid); r++ {
		for c := 0; c < len(universeGrid[0]); c++ {
			if universeGrid[r][c] == "#" {
				galaxies = append(galaxies, Point{x: r, y: c})
			}
		}
	}
	totalDistance := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += manhattanDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Println(totalDistance)
	fmt.Println(galaxies)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
