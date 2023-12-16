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

func warpedManhattanDistance(p1 Point, p2 Point, tsCols []int, tsRows []int) int {
	warpedDistance := int(math.Abs(float64(p1.x-p2.x))) + int(math.Abs(float64(p1.y-p2.y)))
	growthFactor := 1000000 - 1
	for _, sinkCol := range tsCols {
		if (p1.x > sinkCol && p2.x < sinkCol) || (p1.x < sinkCol && p2.x > sinkCol) {
			warpedDistance += growthFactor
		}
	}
	for _, sinkRow := range tsRows {
		if (p1.y > sinkRow && p2.y < sinkRow) || (p1.y < sinkRow && p2.y > sinkRow) {
			warpedDistance += growthFactor
		}
	}
	return warpedDistance
}
func findTimeSinkRows(grid [][]string) []int {
	tsRows := make([]int, 0)
	for r := 0; r < len(grid); r++ {
		if slices.Index(grid[r], "#") == -1 {
			tsRows = append(tsRows, r)
		}
	}
	return tsRows
}

func findTimeSinkCols(grid [][]string) []int {
	tsCols := make([]int, 0)
	for c := 0; c < len(grid[0]); c++ {
		duplicateCol := true
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == "#" {
				duplicateCol = false
			}
		}
		if duplicateCol {
			tsCols = append(tsCols, c)
		}
	}
	return tsCols
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

	timeSinkRows := findTimeSinkRows(universeGrid)
	timeSinkCols := findTimeSinkCols(universeGrid)

	for r := 0; r < len(universeGrid); r++ {
		for c := 0; c < len(universeGrid[0]); c++ {
			if universeGrid[r][c] == "#" {
				galaxies = append(galaxies, Point{x: c, y: r})
			}
		}
	}
	totalDistance := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			totalDistance += warpedManhattanDistance(galaxies[i], galaxies[j], timeSinkCols, timeSinkRows)
		}
	}
	fmt.Println(totalDistance)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
