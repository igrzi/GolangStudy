package functions

import "fmt"

func GetInput() (int, int, int, error) {
	var a, b, c int
	fmt.Print("Digite três números separados por espaço: ")
	_, err := fmt.Scan(&a, &b, &c)
	return a, b, c, err
}
