package main

import "testing"

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	if resultado != esperado {
		t.Errorf("Resultado: %s, esperado: %s", resultado, esperado)
	}
}

func confirmaErro(t *testing.T, resultado error, esperado error) {
	t.Helper()
	if ErroSaldoInsuficiente == nil {
		t.Fatal("Esperava um erro porém não ocorreu")
	}

	if resultado != esperado {
		t.Errorf("Resultado: %s, esperado: %s", resultado, esperado)
	}
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

func TestCarteira(t *testing.T) {

	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(20))
		confirmaSaldo(t, carteira, Bitcoin(20))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}
