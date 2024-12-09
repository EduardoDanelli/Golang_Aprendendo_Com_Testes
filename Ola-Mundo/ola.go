package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const gaucho = "gaucho"
const prefixoOla = "Olá, "
const prefixoHola = "Hola, "
const prefixoFrances = "Bounjour, "
const prefixoGaucho = "Opa meu guri, "

func Ola(nome, idioma string) string {

	if nome == "" {
		nome = "mundo"
	}

	return saudacao(idioma) + nome
}

func saudacao(idioma string) (prefixo string) {

	switch idioma {
	case espanhol:
		prefixo = prefixoHola
	case frances:
		prefixo = prefixoFrances
	case gaucho:
		prefixo = prefixoGaucho
	default:
		prefixo = prefixoOla
	}

	return
}

func main() {
	fmt.Println(Ola("Edu", "português"))
}
