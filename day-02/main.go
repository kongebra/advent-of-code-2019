package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// Opcode type
type Opcode int

const (
	// Add opcode
	Add Opcode = 1
	// Multiply opcode
	Multiply Opcode = 2
	// Halt opcode
	Halt Opcode = 99
)

func main() {

}

func getValuesFromFile(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataString := string(data)
	stringSlice := strings.Split(dataString, ",")

	return convertStringToIntSlice(stringSlice)
}

func convertStringToIntSlice(input []string) ([]int, error) {
	values := make([]int, 0)

	for _, str := range input {
		value, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

// PartOne code
func PartOne() (int, error) {
	values, err := getValuesFromFile("input.txt")
	if err != nil {
		return 0, err
	}

	values[1] = 12
	values[2] = 2

	var current Opcode
	var index = 0

	for current != Halt {
		current = Opcode(values[index])

		first := values[index+1]
		second := values[index+2]
		third := values[index+3]

		if current == Add {
			values[third] = values[first] + values[second]
		} else if current == Multiply {
			values[third] = values[first] * values[second]
		} else if current == Halt {
			break
		} else {
			return 0, errors.New("something went wrong")
		}

		index += 4
	}

	return values[0], nil
}
