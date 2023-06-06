package main

import (
	"fmt"
	"math"
)

func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func preencherNumerosPrimos(ptr *[]int, tamanho int) {
	numPrimos := 0
	numAtual := 2

	for numPrimos < tamanho {
		if isPrimo(numAtual) {
			*ptr = append(*ptr, numAtual)
			numPrimos++
		}
		numAtual++
	}
}

func main() {
	numerosPrimos := []int{}
	preencherNumerosPrimos(&numerosPrimos, 5)
	fmt.Println(numerosPrimos) // Saída: [2 3 5 7 11]
}
