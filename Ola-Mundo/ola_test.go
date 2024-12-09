package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Edu", "")
	esperado := "Olá, Edu"

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}

func TestOla2(t *testing.T) {
	resultado := Ola("Edu!", "")
	esperado := "Olá, Edu!"

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}

func TestOla3(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Edu", "")
		esperado := "Olá, Edu"

		if resultado != esperado {
			verificaMensagemCorreta(t, resultado, esperado)
		}
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		if resultado != esperado {
			verificaMensagemCorreta(t, resultado, esperado)
		}
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Edu", "espanhol")
		esperado := "Hola, Edu"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Edu", "francês")
		esperado := "Bounjour, Edu"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em gaucho", func(t *testing.T) {
		resultado := Ola("Edu", "gaucho")
		esperado := "Opa meu guri, Edu"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
