package main

import (
	"fmt"
	"functions/functions"
)

type Item = functions.Item

func main() {
	items := []Item{} // Array de itens, usando a nova struct
	capacidadeMochila := 0
	for {
		utilidade, peso, quantidade, err := functions.GetInput()
		if err != nil {
			fmt.Println("Erro ao ler os números:", err)
			break
		}

		if utilidade == -1 && peso == -1 && quantidade == -1 {
			fmt.Print("Qual a capacidade da mochila? ")
			fmt.Scan(&capacidadeMochila)
			break
		}

		calculo := float64(utilidade) / float64(peso)

		// Criar uma instância da struct Item e adicionar ao array
		items = append(items, Item{Value: calculo, OriginalIndex: len(items), Peso: peso, Quantidade: quantidade})
	}

	sortedItems := functions.SelectionSortDecrescente(items)

	currentWeight := 0
	for _, item := range sortedItems {
		if currentWeight+item.Peso <= capacidadeMochila && item.Quantidade > 0 {
			quantidadeToAdd := min(item.Quantidade, (capacidadeMochila-currentWeight)/item.Peso)
			if quantidadeToAdd > 0 {
				fmt.Printf("%d %d\n", item.OriginalIndex, quantidadeToAdd)
				currentWeight += quantidadeToAdd * item.Peso
			}
		}
	}
}
