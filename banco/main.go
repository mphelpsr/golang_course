package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaMarcelo := contas.ContaPoupanca{}

	contaMarcelo.DepositarValor(500)
	contaMarcelo.SacarValores(100)

	fmt.Println(contaMarcelo)
	fmt.Println(contaMarcelo.Obtersaldo())

}
