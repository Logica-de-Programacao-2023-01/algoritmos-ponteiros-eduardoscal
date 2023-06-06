package main

import "fmt"

type LivroComDesconto struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(ptr *LivroComDesconto) {
	ptr.Preco = ptr.Preco * 0.9
}

func main() {
	livro := LivroComDesconto{
		Titulo: "Livro A",
		Autor:  "Autor B",
		Preco:  50.0,
	}
	aplicarDesconto(&livro)
	fmt.Println(livro.Preco)
}
