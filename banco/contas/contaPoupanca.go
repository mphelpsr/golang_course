package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) SacarValores(valor float64) (string, float64) {

	if valor < 0 && valor >= c.saldo {
		c.saldo -= valor
	} else {
		return "Seu saque nao foi afetado! saldo atual: ", c.saldo
	}

	return "Valor sacado com sucesso! Seu novo saldo é: ", c.saldo
}

func (c *ContaPoupanca) DepositarValor(valor float64) (string, float64) {

	if valor > 0 {
		c.saldo += valor
	} else {
		return "Valor nao foi depositado! saldo atual: ", c.saldo
	}

	return "Valor depositado com sucesso! Seu novo saldo é: ", c.saldo
}

func (c *ContaPoupanca) Obtersaldo() float64 {
	return c.saldo
}
