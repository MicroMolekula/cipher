package vizhener

import "strings"

func getShiftsArray(word string) []int {
	wordLower := strings.ToLower(word)
	alphaF, _ := getAlphabet(wordLower)
	chars := []rune(wordLower)
	result := make([]int, len(chars))
	for i := 0; i < len(chars); i++ {
		result[i] = int(chars[i]) - int(alphaF)
	}
	return result
}

func getAlphabet(word string) (firstChar, lastChar rune) {
	char := []rune(word)[0]
	russianChar := 'а'
	englishChar := 'a'
	switch {
	case char >= 'a' && char <= 'z':
		firstChar = englishChar
		lastChar = 'z'
	case char >= 'а' && char <= 'я':
		firstChar = russianChar
		lastChar = 'я'
	}
	return
}

func Decode(word, key string) string {
	shiftsArray := getShiftsArray(key)
	wordChars := []rune(word)
	result := make([]rune, len(wordChars))
	for i := 0; i < len(wordChars); i++ {
		cipherIndex := i % len(shiftsArray)
		result[i] = getWordChar(wordChars[i], shiftsArray[cipherIndex])
	}
	return string(result)
}

func getWordChar(char rune, shift int) rune {
	alphaF, alphaL := getAlphabet(string(char))
	if int(char)-shift < int(alphaF) {
		return rune(int(alphaL) - (int(alphaF) - (int(char) - shift)) + 1)
	}
	return rune(int(char) - shift)
}

func Code(word, key string) string {
	shiftsArray := getShiftsArray(key)
	wordChars := []rune(word)
	result := make([]rune, len(wordChars))
	for i := 0; i < len(wordChars); i++ {
		cipherIndex := i % len(shiftsArray)
		result[i] = getCipherChar(wordChars[i], shiftsArray[cipherIndex])
	}
	return string(result)
}

func getCipherChar(char rune, shift int) rune {
	alphaF, alphaL := getAlphabet(string(char))
	if int(char)+shift > int(alphaL) {
		return rune(((int(char) + shift) - int(alphaL)) + int(alphaF) - 1)
	}
	return rune(int(char) + shift)
}
