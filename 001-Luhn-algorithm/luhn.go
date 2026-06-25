package main

import (
	"fmt"
)

func LuhnValid(cardNumber string) bool {
	if len(cardNumber) < 2 {
		return false
	}

	sum := 0
	alt := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		ch := cardNumber[i]
		if ch < '0' || ch > '9' {
			continue
		}

		d := int(ch - '0')

		if alt {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
		alt = !alt
	}

	return sum%10 == 0
}

// Pre-computed lookup table for doubled digits
var doubleLUT = [10]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

func LuhnValidLUT(cardNumber string) bool {
	if len(cardNumber) < 2 {
		return false
	}

	sum := 0
	alt := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		ch := cardNumber[i]
		if ch < '0' || ch > '9' {
			continue
		}

		d := int(ch - '0')

		if alt {
			sum += doubleLUT[d]
		} else {
			sum += d
		}
		alt = !alt
	}

	return sum%10 == 0
}

func main() {
	LuhnValid("17893729974")

	fmt.Println("ok")
}
