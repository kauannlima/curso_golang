package main

import (
	"exercicios/model"
	"fmt"
	"time"
)

func main() {
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "Arroz", "Feijão", "Macarrão")
	var compra = model.NewCompra("Mercado XYZ", time.Now(), nomeDosItens)

	fmt.Println(compra)
}
