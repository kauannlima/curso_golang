package main

import (
	"fmt"
	"golangestudos/model"
)

func main() {
	fmt.Println("Iniciando...")

	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "XYZ-1234",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 160,
	}

	fmt.Println("Moto:", moto)

}
