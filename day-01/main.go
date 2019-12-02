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

}

// FuelCounterUpper calulates fuel needed based on mass
func FuelCounterUpper(mass float64) float64 {
	// Divide mass by 3
	result := mass / 3.0
	// Round down
	result = math.Floor(result)
	// Subtract 2
	result -= 2

	return result
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

	// Print fuel requirements
	fmt.Printf("Fuel requirements: %.0f", sum)
	
	// Check for any scanning errors, and handle them
	if err := scanner.Err(); err != nil {
		return 0.0, err
	}

	return sum, nil
}
