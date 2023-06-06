package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValorAoSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}
