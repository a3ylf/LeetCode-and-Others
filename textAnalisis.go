//Inspecting a text for the amount of words, the most frequent word and a map of the appearance of words

package main

import "strings"
import "fmt"

func main() {
	texto := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut vitae nunc id quam fringilla efficitur. Sed vitae varius urna. Integer quis orci sit amet arcu efficitur sollicitudin. Duis auctor arcu ut est viverra mollis. Sed vel arcu vel lorem ullamcorper viverra. Curabitur commodo quam non leo aliquam malesuada. Ut non neque vitae lacus fermentum accumsan. Duis hendrerit massa vel enim tincidunt auctor. Vivamus eget metus lacinia, eleifend tortor sit amet, pellentesque mi.`
	wordAmounts, wordCounter, mostFrequent := textAnalisis(texto)
	fmt.Printf("Amount of words: %v\nMost frequent word: %v\n\n", wordAmounts, mostFrequent)
	fmt.Println(wordCounter)
}
func textAnalisis(texto string) (wordAmounts int, counter map[string]int, mostFrequent string) {
	counter = make(map[string]int)
	mostFrequent = "<>><>><>><><>><><><>><<>>"
	palavras := strings.Fields(strings.ToLower(texto))
	wordAmounts = len(palavras)
	for _, palavra := range palavras {
		counter[palavra]++
		if counter[palavra] > counter[mostFrequent] {
			mostFrequent = palavra
		}
	}
	return wordAmounts, counter, mostFrequent
}
