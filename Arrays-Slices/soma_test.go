package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	// t.Run("array de 5 números", func(t *testing.T) {
	// 	numeros := []int{1, 2, 3, 4, 5}

	// 	resultado := Soma(numeros)
	// 	esperado := 15

	// 	if resultado != esperado {
	// 		t.Errorf("Resultado: %d, esperado: %d, dados: %v", resultado, esperado, numeros)
	// 	}
	// })

	t.Run("usando slice(array sem tamanho definido)", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		resultado := Soma(numeros)
		esperado := 45

		if resultado != esperado {
			t.Errorf("Resultado: %d, esperado: %d, dados: %v", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado: %d, esperado: %d", resultado, esperado)
	}
}

func TestSomaTodoResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("Resultado: %d, esperado: %d", resultado, esperado)
		}
	}

	t.Run("soma os slices", func(t *testing.T) {
		resultado := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma os slices", func(t *testing.T) {
		resultado := SomaTodoResto([]int{}, []int{0, 9})
		esperado := []int{0, 9}
		verificaSomas(t, resultado, esperado)
	})
}
