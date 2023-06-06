package main

import "fmt"

func somaDosUltimosDigitos(ptr *int) {
	lastDigit := *ptr % 10
	*ptr /= 10
	secondLastDigit := *ptr % 10
	*ptr = lastDigit + secondLastDigit
}
