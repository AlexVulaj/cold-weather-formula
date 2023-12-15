package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	temperature := promptUser(reader, "What's the current temperature outside in fahrenheit?", "temperature")
	numDrinks := promptUser(reader, "How many drinks have you had?", "number of drinks")
	adjustedTotal := calculateAdjustedTotal(temperature, numDrinks)

	fmt.Printf("Your adjusted winter drinking total is: %v", adjustedTotal)
}

func promptUser(reader *bufio.Reader, prompt string, inputType string) int {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	var result int
	for err := errors.New(""); err != nil; {
		result, err = strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%v is not a valid %v.\n", input, inputType)
			return promptUser(reader, prompt, inputType)
		}
	}
	return result
}

func calculateAdjustedTotal(temp int, numDrinks int) int {
	return int((1 - (float32(70-temp) / 100)) * float32(numDrinks)) // always round down
}
