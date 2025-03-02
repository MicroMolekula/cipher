package lr

import (
	"fmt"
	"github.com/cipher/gamification"
	"github.com/cipher/vizhener"
	"os"
	"strings"
)

func Lr1() {
	fmt.Println("Программа для шифрования и дешифрования")
	fmt.Println("_______________________________________")
	for {
		var method int
		fmt.Print("Выберите метод шифрования:\n1 - Виженера\n2 - Гаммирование\n> ")
		fmt.Scan(&method)
		switch method {
		case 1:
			var fileOrKeyboard int
			fmt.Print("1 - Из файла\n2 - С клавиатуры\n> ")
			fmt.Scan(&fileOrKeyboard)
			switch fileOrKeyboard {
			case 1:
				FileVizhiner()
			case 2:
				Vizhiner()
			}
		case 2:
			var fileOrKeyboard int
			fmt.Print("1 - Из файла\n2 - С клавиатуры\n> ")
			fmt.Scan(&fileOrKeyboard)
			switch fileOrKeyboard {
			case 1:
			case 2:
				Gamification()
			}
		default:
			break
		}
	}
}

func Vizhiner() {
	var flag int
	fmt.Print("\n1 - Шифровать\n2-Дешифровать\n> ")
	fmt.Scan(&flag)
	switch flag {
	case 1:
		var word string
		var key string
		fmt.Print("Слово для шифрования: ")
		fmt.Scanf("%s", &word)
		fmt.Print("Ключ: ")
		fmt.Scanf("%s", &key)
		result := vizhener.Code(word, key)
		fmt.Println("Результат шифрования: ", result)
	case 2:
		var word string
		var key string
		fmt.Print("Слово для дешифрования: ")
		fmt.Scanf("%s", &word)
		fmt.Print("Ключ: ")
		fmt.Scanf("%s", &key)
		result := vizhener.Decode(word, key)
		fmt.Println("Результат дешифрования: ", result)
	}
}

func Gamification() {
	var flag int
	fmt.Print("\n1 - Шифровать\n2 - Дешифровать\n> ")
	fmt.Scan(&flag)
	switch flag {
	case 1:
		var word string
		fmt.Print("Слово для шифрования: ")
		fmt.Scanf("%s", &word)
		result, cipher := gamification.Code(word)
		fmt.Println("Результат шифрования: ", result)
		fmt.Println("Ключ шифрования: ", cipher)
	case 2:
		var word string
		var key string
		fmt.Print("Слово для дешифрования: ")
		fmt.Scanf("%s", &word)
		fmt.Print("Ключ: ")
		fmt.Scanf("%s", &key)
		result := gamification.Decode(word, key)
		fmt.Println("Результат дешифрования: ", result)
	}
}

func FileVizhiner() {
	var file string
	var key string
	fmt.Print("Файл: ")
	fmt.Scanf("%s", &file)
	fmt.Print("Ключ: ")
	fmt.Scanf("%s", &key)
	fileData, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(fileData), " ")
	result := ""
	var flag int
	fmt.Print("\n1 - Шифровать\n2-Дешифровать\n> ")
	fmt.Scan(&flag)
	switch flag {
	case 1:
		for _, word := range words {
			result += vizhener.Code(word, key) + " "
		}
		fmt.Println("Результат шифрования: ", result)
	case 2:
		for _, word := range words {
			result += vizhener.Decode(word, key) + " "
		}
		fmt.Println("Результат шифрования: ", result)
	}
}
