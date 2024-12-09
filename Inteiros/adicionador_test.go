package main

import (
	"testing"
)

func TestAdicionador(t *testing.T) {
	resultado := Adiciona(2, 2)
	esperado := 4

	if resultado != esperado {
		t.Errorf("Resultado: '%d', esperado: '%d'", resultado, esperado)
	}
}
