package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) transferenciaEntreContas(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.DepositarValor(valorTransferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) SacarValores(valor float64) (string, float64) {

	if valor < 0 && valor >= c.Saldo {
		c.Saldo -= valor
	} else {
		return "Seu saque nao foi afetado! Saldo atual: ", c.Saldo
	}

	return "Valor sacado com sucesso! Seu novo Saldo é: ", c.Saldo
}

func (c *ContaCorrente) DepositarValor(valor float64) (string, float64) {

	if valor > 0 {
		c.Saldo += valor
	} else {
		return "Valor nao foi depositado! Saldo atual: ", c.Saldo
	}

	return "Valor depositado com sucesso! Seu novo Saldo é: ", c.Saldo
}
