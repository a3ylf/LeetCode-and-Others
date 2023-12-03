//Mapeando a quantidade de aparições de um nome com base na primeira letra desse nome

package main

import (
	"fmt"
)

func main() {
	nomes := []string{"João", "João", "Joaquim", "Bruno", "João", "João", "João", "João", "João", "Bruno", "Carlos"}
	mapa := nameCounter(nomes)
	for letter, insideMap := range mapa {
		fmt.Printf("\n\nLetra: %s", string(letter))
		for names, ammount := range insideMap {
			fmt.Printf("\n	Nome: %v, Quantidade: %v", names, ammount)
		}
	}
}

func nameCounter(names []string) map[rune]map[string]int {

	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if name == "" {
			continue
		}

		firstchar := rune(name[0])
		_, ok := counts[firstchar]
		if !ok {
			counts[firstchar] = make(map[string]int)
		}
		counts[firstchar][name]++

	}
	return counts
}
