package main

import (
	"regexp"
	"strconv"
)

func main() {
	firstDigitConstant := 100
	inputStart := 372037
	inputEnd := 905157
	passwords := make([]int, 0)

	for number := inputStart; number <= inputEnd; number++ {
		// validate repeating adjacent digit
		isSixDigitsWithRepeating, _ := regexp.MatchString("(\\d){0,6}\\1", strconv.Itoa(number))
		if isSixDigitsWithRepeating {
			continue
		}

		// validate decreasing/same digits
		decreases := false
		prevDigit := firstDigitConstant
		for _, digit := range [...]int{number} {
			if prevDigit == firstDigitConstant {
				prevDigit = digit
				continue
			} else {
				if prevDigit > digit {
					decreases = true
					break
				}
			}
			prevDigit = digit
		}

		if !decreases {
			passwords = append(passwords, number)
		}
	}

	println("valid passwords: " + strconv.Itoa(len(passwords)))
}
