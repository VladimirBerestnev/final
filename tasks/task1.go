package tasks

import (
	"strings"
)

// Напишите программу на языке Go, которая проверяет, является ли введенная строка палиндромом.

func Polindrome(s string) bool {

	noSpace := strings.ReplaceAll(s, " ", "")
	noDot := strings.ReplaceAll(noSpace, ".", "")
	noComma := strings.ReplaceAll(noDot, ",", "")
	final := strings.ToLower(noComma)
	reverse := ""

	for i := len(final) - 1; i > -1; i-- {
		reverse = reverse + (string(final[i]))
	}

	if final == reverse {
		return true
	} else {
		return false
	}
}
