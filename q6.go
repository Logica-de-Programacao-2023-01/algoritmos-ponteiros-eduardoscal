package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func alterarTituloSeAutorAnonimo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}
