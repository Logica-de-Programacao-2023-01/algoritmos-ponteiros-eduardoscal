package main

import "fmt"

func modificarVariavel(ptr *int) {
	*ptr = 42
}

func main() {
	numero := 10
	modificarVariavel(&numero)
	fmt.Println(numero)
}
