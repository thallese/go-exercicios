/*
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
  - Crie um tipo para um struct chamado "pessoa"
  - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
  - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
  - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
  - Demonstre no seu código:
  - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
  - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"

- Se precisar de dicas, veja: https://gobyexample.com/interfaces
- Solução: https://github.com/ellenkorbes/aprend...
*/
package main

import (
	"fmt"
	"strconv"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p *pessoa) falar() string {
	frase := "Olá, eu sou " + p.Nome + " " + p.Sobrenome + " e tenho " + strconv.Itoa(p.Idade) + " anos"
	return frase
}

type humanos interface {
	falar() string
}

func main() {
	thalles := pessoa{
		Nome:      "Thalles",
		Sobrenome: "Éboli",
		Idade:     26,
	}

	dizerAlgumaCoisa(&thalles)
}

func dizerAlgumaCoisa(h humanos) {
	fmt.Println(h.falar())
	fmt.Printf("%T", h)
}
