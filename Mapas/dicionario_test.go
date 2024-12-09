package main

import "testing"

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s, dado: %s", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s", resultado, esperado)
	}
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {

	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada: ", err)
	}

	if definicao != resultado {
		t.Errorf("resultado: %s, esperado: %s", resultado, definicao)
	}
}

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isto é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isto é apenas um teste"
		comparaString(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {

	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isto é apenas um teste"

		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isto é apenas um teste"
		dicionario := Dicionario{palavra: definicao}

		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isto é apenas um teste"
		novaDefinicao := "nova definição"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isto é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)
		comparaErro(t, err, ErrPalavraInexistente)
	})
}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s', seja deletado", palavra)
	}
}
