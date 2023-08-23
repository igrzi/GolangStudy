package main

import (
	"fmt"
)

type Item struct {
	Value         float64 // Valor calculado de utilidade por peso
	OriginalIndex int     // Índice original do item antes de ordenação
	Peso          int     // Peso do item
	Quantidade    int     // Quantidade do item
}

func selectionSortDecrescente(arr []Item) []Item {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			// Tratar empates considerando o índice original
			if arr[j].Value == arr[maxIndex].Value {
				if arr[j].OriginalIndex < arr[maxIndex].OriginalIndex {
					maxIndex = j
				}
			} else if arr[j].Value > arr[maxIndex].Value {
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
	return arr
}

func getInput() (int, int, int, error) {
	var a, b, c int
	fmt.Print("Digite três números separados por espaço: ")
	_, err := fmt.Scan(&a, &b, &c)
	return a, b, c, err
}

func main() {
	items := []Item{} // Array de itens, usando a nova struct
	capacidadeMochila := 0
	for {
		utilidade, peso, quantidade, err := getInput()
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

	sortedItems := selectionSortDecrescente(items)

	for _, item := range sortedItems {
		originalIndex := item.OriginalIndex
		peso := item.Peso
		quantidade := item.Quantidade
		fmt.Printf("Índice: %d, Peso: %d, Quantidade: %d\n", originalIndex, peso, quantidade)
	}
}
