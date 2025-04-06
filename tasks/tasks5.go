package tasks

import "unicode"

// Дана произвольная строка цифр и букв. Необходимо вырезать буквы и вывести результат.

func GetLetters(s string) string {
	letters := ""
	for _, char := range s {
		if !unicode.IsNumber(char) {
			letters = letters + string(char)
		}
	}
	return letters
}
