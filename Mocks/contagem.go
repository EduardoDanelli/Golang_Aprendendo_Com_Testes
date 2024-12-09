package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Vai"
const inicioContagem = 3

type Sleeper interface {
	Sleeper()
}

type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleeper() {
	s.Chamadas++
}

type SleeperPadrao struct {
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

const escrita = "escrita"
const pausa = "pausa"

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

func (c *SleeperConfiguravel) Pausa() {

}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
