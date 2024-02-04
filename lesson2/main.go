package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	result := map[rune]int{}
	coutLetters := 0
	// Вставьте ваш код здесь
	for _, r := range text {
		if unicode.IsLetter(r) {
			result[unicode.ToLower(r)]++
			coutLetters++
		}
	}

	for k, v := range result {
		percentage := float64(v) / float64(coutLetters) * 100
		fmt.Printf("%c: %d (%.2f%%)\n", k, v, percentage)
	}
}
