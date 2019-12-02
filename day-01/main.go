package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	result, err := PartTwo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.0f", result)
}

// FuelCounterUpper calulates fuel needed based on mass
func FuelCounterUpper(mass float64) float64 {
	return math.Floor(mass/3.0) - 2
}

// PartOne code
func PartOne() (float64, error) {
	// Open file, and check for any errors
	file, err := os.Open("input.txt")
	if err != nil {
		return 0.0, err
	}

	// Close the file after we use it :)
	defer file.Close()

	// Create a new scanner from the file
	scanner := bufio.NewScanner(file)

	sum := 0.0

	// Keep on scanning the file line by line
	for scanner.Scan() {
		// Get the text of the line
		text := scanner.Text()

		// Convert string to float64, and check for error
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return 0.0, err
		}

		// Calculate fuel, and add it to the sum
		sum += FuelCounterUpper(value)
	}

	// Check for any scanning errors, and handle them
	if err := scanner.Err(); err != nil {
		return 0.0, err
	}

	return sum, nil
}

// PartTwo code
func PartTwo() (float64, error) {
	// Open file, and check for any errors
	file, err := os.Open("input.txt")
	if err != nil {
		return 0.0, err
	}

	// Close the file after we use it :)
	defer file.Close()

	// Create a new scanner from the file
	scanner := bufio.NewScanner(file)

	sum := 0.0

	// Keep on scanning the file line by line
	for scanner.Scan() {
		// Get the text of the line
		text := scanner.Text()

		// Convert string to float64, and check for error
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return 0.0, err
		}

		// Calculate fuel, and add it to the sum
		sum += PartTwoHelper(value)
	}

	// Check for any scanning errors, and handle them
	if err := scanner.Err(); err != nil {
		return 0.0, err
	}

	return sum, nil
}

// PartTwoHelper code
func PartTwoHelper(value float64) float64 {
	// Set result to 0
	result := 0.0
	// Store value in temporary variable
	temp := value
	// Loop while temp is larger than 0
	for FuelCounterUpper(temp) > 0 {
		// Set temp
		temp = FuelCounterUpper(temp)
		// Add new temp value to result
		result += temp
	}
	// Return result
	return result
}
