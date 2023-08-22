package main

import (
	"fmt"
)

type Item struct {
	Value         float64 // Valor calculado de utilidade por peso
	OriginalIndex int     // Índice original do item antes de ordenação
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

func utilidadePeso(a, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	items := []Item{} // Array de itens, usando a nova struct

	for {
		utilidade, peso, quantidade, err := getInput()
		if err != nil {
			fmt.Println("Erro ao ler os números:", err)
			break
		}

		if utilidade == -1 && peso == -1 && quantidade == -1 {
			break
		}

		calculo := utilidadePeso(utilidade, peso)

		// Criar uma instância da struct Item e adicionar ao array
		items = append(items, Item{Value: calculo, OriginalIndex: len(items)})
	}

	sortedItems := selectionSortDecrescente(items)

	// Imprimir os índices originais dos itens após ordenação
	for _, item := range sortedItems {
		originalIndex := item.OriginalIndex
		value := item.Value
		fmt.Printf("Índice Original: %d, Valor: %.2f\n", originalIndex, value)
	}
}
