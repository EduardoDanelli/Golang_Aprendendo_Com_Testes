package main

import (
	"bytes"
	"testing"
)

func TestCumprimento(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Edu")

	resultado := buffer.String()
	esperado := "Olá, Edu"

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s", resultado, esperado)
	}
}
