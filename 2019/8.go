package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := reader.ReadString('\n')

	var digits []string
	for i := 0; i < len(input); i++ {
		digits = append(digits, string(input[i]))
	}
	digitsCopy := digits

	imageWidth := 25
	imageHeight := 6
	totalDigitsInImage := imageHeight * imageWidth
	layers := 0
	var numOfZerosInLayers []int

	for len(digits) > 0 {
		numOfZerosInLayer := 0

		// remove each layer, count zeros
		index := 0
		for i := 0; i < totalDigitsInImage; i++ {
			x := digits[index]        // get the 0 index element from slice
			digits = digits[index+1:] // remove the 0 index element from slice
			if x == "0" {
				numOfZerosInLayer++
			}
			index++
		}

		numOfZerosInLayers = append(numOfZerosInLayers, int(numOfZerosInLayer))
	}
}
