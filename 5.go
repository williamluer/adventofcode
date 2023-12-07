package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func convertToInts(s []string) []int {
	output := make([]int, 0)
	for i := 0; i < len(s); i += 1 {
		val, _ := strconv.Atoi(s[i])
		output = append(output, val)
	}
	return output
}

func getVal(inputSlice []MapRow, currentVal int) int {
	var output int
	var closestRow MapRow
	// fmt.Println("currentVal ", currentVal)
	// fmt.Println("maprow ", inputSlice)
	for i := 0; i < len(inputSlice); i++ {
		if i == len(inputSlice)-1 && inputSlice[i].sourceStart < currentVal {
			// This case happens when we are at the end of the slice
			// it could be within range or it could not be
			closestRow = inputSlice[i]
		} else if i == 0 && inputSlice[i].sourceStart > currentVal {
			// This happens if currentVal is less than the first element.
			// Always return self
			return currentVal
		} else if i == 0 && inputSlice[i].sourceStart == currentVal {
			return inputSlice[i].destinationStart
		} else if inputSlice[i].sourceStart > currentVal {
			// this happens if it's in the middle
			closestRow = inputSlice[i-1]
			break
		} else if inputSlice[i].sourceStart == currentVal {
			// This happens if it's equal
			return inputSlice[i].destinationStart
		}
	}
	if closestRow.sourceStart+closestRow.rangeLenth >= currentVal {
		output = closestRow.destinationStart + (currentVal - closestRow.sourceStart)
	} else {
		output = currentVal
	}
	// fmt.Println("Returning ", output)
	// fmt.Println()
	return output
}

type MapRow struct {
	destinationStart int
	sourceStart      int
	rangeLenth       int
}

func main() {
	filename := "5.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reg := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(file)
	i := 0
	seeds := make([]int, 0)
	seedToSoil := make([]MapRow, 0)            // 3
	soilToFertilizer := make([]MapRow, 0)      // 38
	fertilizerToWater := make([]MapRow, 0)     // 64
	waterToLight := make([]MapRow, 0)          // 88
	lightToTemperature := make([]MapRow, 0)    //109
	temperatureToHumidity := make([]MapRow, 0) //122
	humidityToLocation := make([]MapRow, 0)
	fmt.Println("Creating data")

	for scanner.Scan() {
		current_line := scanner.Text()
		lineVals := reg.FindAllString(current_line, -1)
		if len(lineVals) == 0 {
			i += 1
			continue
		}
		vals := convertToInts(lineVals)
		currentRow := MapRow{destinationStart: vals[0], sourceStart: vals[1], rangeLenth: vals[2]}
		if i == 0 {
			seeds = convertToInts(lineVals)
		} else if i > 132 {
			humidityToLocation = append(humidityToLocation, currentRow)
		} else if i > 121 {
			temperatureToHumidity = append(temperatureToHumidity, currentRow)
		} else if i > 108 {
			lightToTemperature = append(lightToTemperature, currentRow)
		} else if i > 87 {
			waterToLight = append(waterToLight, currentRow)
		} else if i > 63 {
			fertilizerToWater = append(fertilizerToWater, currentRow)
		} else if i > 37 {
			soilToFertilizer = append(soilToFertilizer, currentRow)
		} else if i > 2 {
			seedToSoil = append(seedToSoil, currentRow)
		}

		// if i == 0 {
		// 	seeds = vals
		// } else if i > 30 {
		// 	humidityToLocation = append(humidityToLocation, currentRow)
		// } else if i > 26 {
		// 	temperatureToHumidity = append(temperatureToHumidity, currentRow)
		// } else if i > 21 {
		// 	lightToTemperature = append(lightToTemperature, currentRow)
		// } else if i > 17 {
		// 	waterToLight = append(waterToLight, currentRow)
		// } else if i > 11 {
		// 	fertilizerToWater = append(fertilizerToWater, currentRow)
		// } else if i > 6 {
		// 	soilToFertilizer = append(soilToFertilizer, currentRow)
		// } else if i > 2 {
		// 	seedToSoil = append(seedToSoil, currentRow)
		// }
		i += 1
	}

	fmt.Println("Sorting...")
	sort.Slice(seedToSoil, func(i, j int) bool {
		return seedToSoil[i].sourceStart < seedToSoil[j].sourceStart
	})
	sort.Slice(soilToFertilizer, func(i, j int) bool {
		return soilToFertilizer[i].sourceStart < soilToFertilizer[j].sourceStart
	})
	sort.Slice(fertilizerToWater, func(i, j int) bool {
		return fertilizerToWater[i].sourceStart < fertilizerToWater[j].sourceStart
	})
	sort.Slice(waterToLight, func(i, j int) bool {
		return waterToLight[i].sourceStart < waterToLight[j].sourceStart
	})
	sort.Slice(lightToTemperature, func(i, j int) bool {
		return lightToTemperature[i].sourceStart < lightToTemperature[j].sourceStart
	})
	sort.Slice(temperatureToHumidity, func(i, j int) bool {
		return temperatureToHumidity[i].sourceStart < temperatureToHumidity[j].sourceStart
	})
	sort.Slice(humidityToLocation, func(i, j int) bool {
		return humidityToLocation[i].sourceStart < humidityToLocation[j].sourceStart
	})

	fmt.Println("HERE ARE THE SORTED OUTPUTS")
	fmt.Println(seedToSoil)
	fmt.Println(soilToFertilizer)
	fmt.Println(fertilizerToWater)
	fmt.Println(waterToLight)
	fmt.Println(lightToTemperature)
	fmt.Println(temperatureToHumidity)
	fmt.Println(humidityToLocation)

	fmt.Println()
	minLocation := 999999999999999999
	fmt.Println("Iterating over seeds")
	for _, seed := range seeds {
		// fmt.Println("\n\n\nON seed number ", k)
		soil := getVal(seedToSoil, seed)

		// fmt.Println("soil ", soil)
		fertilizer := getVal(soilToFertilizer, soil)

		// fmt.Println("ertilizer ", fertilizer)
		water := getVal(fertilizerToWater, fertilizer)

		// fmt.Println("water ", water)
		light := getVal(waterToLight, water)

		// fmt.Println("light ", light)
		temperature := getVal(lightToTemperature, light)

		// fmt.Println("temperature ", temperature)
		humidity := getVal(temperatureToHumidity, temperature)

		// fmt.Println("humidity ", humidity)
		location := getVal(humidityToLocation, humidity)

		fmt.Println("location ", location)
		if location < minLocation {
			fmt.Println("Setting min locaiton to ", location)
			minLocation = location
		}
	}

	fmt.Println(minLocation)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
