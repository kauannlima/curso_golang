package main

import (
	"fmt"
	"golangestudos/model"
	"time"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua  29",
		Numero: 2556,
		Cidade: "Santos",
	}
	pessoa := model.Pessoa{
		Nome:             "Joaquim Teste",
		Endereco:         endereco,
		DataDeNascimento: time.Date(2002, 5, 25, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	pessoa.CalculaIdade()
	fmt.Printf("A idade de %s é %d anos\n", pessoa.Nome, pessoa.Idade)
}
