package model

import (
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) *Compra {
	var itens []ItemDaCompra
	for _, nome := range nomeDosItens {
		itens = append(itens, ItemDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    data,
		Itens:   itens,
	}

}
