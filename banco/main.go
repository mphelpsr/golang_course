package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteMarcelo := clientes.Titular{}
	clienteMarcelo.Nome = "Marcelo"
	clienteMarcelo.CPF = "06237654461"
	clienteMarcelo.NomeProfissao = "Analista Sistemas"
	contaMarcelo := contas.ContaCorrente{clienteMarcelo, 222, 12121, 1000.}

	fmt.Println(contaMarcelo)

}
