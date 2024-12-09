package iteracao

const quantRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}
