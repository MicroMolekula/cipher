package gamification

import (
	"math"
	"math/rand"
	"time"
)

const (
	N = 25
)

func getRune(num int) rune {
	return rune('a' + num)
}

func getNum(char rune) int {
	return int(char) - 'a'
}

func generateCipher(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randomizer.Intn(1000000000)
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		seed = int(math.Pow(7, 5)*float64(seed)) % int(math.Pow(2, 31)-1.0)
		result[i] = getRune(seed % N)
	}

	return string(result)
}

func Code(word string) (resultCipher string, cipherWord string) {
	charsWord := []rune(word)
	cipherWord = generateCipher(len(charsWord))
	charsCipher := []rune(cipherWord)
	result := make([]rune, len(charsWord))
	for i := 0; i < len(charsWord); i++ {
		t := getNum(charsWord[i]) + getNum(charsCipher[i])
		if t > N {
			t %= N
		}
		result[i] = getRune(t)
	}
	return string(result), string(charsCipher)
}

func Decode(word, key string) string {
	charsWord := []rune(word)
	charsCipher := []rune(key)
	result := make([]rune, len(charsWord))
	for i := 0; i < len(charsWord); i++ {
		t := getNum(charsWord[i]) - getNum(charsCipher[i])
		if t < 0 {
			t = N + t
		}
		result[i] = getRune(t)
	}
	return string(result)
}
