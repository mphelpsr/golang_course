package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) transferenciaEntreContas(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.DepositarValor(valorTransferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) SacarValores(valor float64) (string, float64) {

	if valor < 0 && valor >= c.saldo {
		c.saldo -= valor
	} else {
		return "Seu saque nao foi afetado! saldo atual: ", c.saldo
	}

	return "Valor sacado com sucesso! Seu novo saldo é: ", c.saldo
}

func (c *ContaCorrente) DepositarValor(valor float64) (string, float64) {

	if valor > 0 {
		c.saldo += valor
	} else {
		return "Valor nao foi depositado! saldo atual: ", c.saldo
	}

	return "Valor depositado com sucesso! Seu novo saldo é: ", c.saldo
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}
