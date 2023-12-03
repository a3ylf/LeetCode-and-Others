package Leetcode

import "strings"

func romanToInt(s string) int {
	var sum int

	s = strings.NewReplacer("CM", "m", "CD", "d", "XL", "l", "XC", "c", "IV", "v", "IX", "x").Replace(s)
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'v': 4,
		'X': 10,
		'x': 9,
		'L': 50,
		'l': 40,
		'C': 100,
		'c': 90,
		'D': 500,
		'd': 400,
		'M': 1000,
		'm': 900,
	}
	for _, letter := range s {

		sum += m[letter]
	}
	return sum
}
