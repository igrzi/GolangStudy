package functions

type Item struct {
	Value         float64 // Valor calculado de utilidade por peso
	OriginalIndex int     // Índice original do item antes de ordenação
	Peso          int     // Peso do item
	Quantidade    int     // Quantidade do item
}

func SelectionSortDecrescente(arr []Item) []Item {
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
