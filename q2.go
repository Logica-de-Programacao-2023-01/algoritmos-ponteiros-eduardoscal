package main

import "fmt"

func verificarParOuImpar(ptr *int) {
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
}

