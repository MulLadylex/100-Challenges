package main

import "strings"

func GetCount1(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

// 通过` strings `的` Count `或` Contains `方法实现：
func GetCount2(strn string) int {
	count := 0

	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)
	}

	return count
}

func countVowels(s string) int {
	vowels := "aeiou"
	count := 0

	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}

	return count
}
