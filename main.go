package main

import (
	"fmt"
	"golangestudos/model"
)

func main() {
	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua  7",
		Numero: 88,
		Cidade: "Santos",
		Estado: "SP",
	}

	fmt.Println(endereco)
	endereco.Numero = 100
	fmt.Println(endereco.Numero)
}
