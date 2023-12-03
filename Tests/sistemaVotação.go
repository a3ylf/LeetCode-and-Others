// Receiving a map of a candidate number and the amount of votes received
//and returning a new map with their names instead

package main

import (
	"fmt"
)

func main() {
	votos := map[int]int{
		1: 37,
		2: 31,
		3: 12,
		4: 26,
	}
	fmt.Println(votacSistema(votos))
}

func votacSistema(votados map[int]int) map[string]int {
	candidates := []string{"Louis", "Bruno", "Joana", "Clara"}
	counter := make(map[string]int)
	for candidate, amount := range votados {
		counter[candidates[candidate-1]] = amount
	}
	return counter
}
