package verify

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
)

func GetHashValueFile(filename string) (*big.Int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}
	hash := sha256.Sum256(data)
	hashInt := new(big.Int).SetBytes(hash[:])
	return hashInt, nil
}

func SaveKeys(privateKey, publicKey, directory string) {
	if err := os.WriteFile(directory+"/publicKey", []byte(publicKey), 0666); err != nil {
		panic(err)
	}
	if err := os.WriteFile(directory+"/privateKey", []byte(privateKey), 0666); err != nil {
		panic(err)
	}
}

func GetKeys(directory string) (string, string, error) {
	publicKey, err := os.ReadFile(directory + "/privateKey")
	if err != nil {
		return "", "", err
	}
	privateKey, err := os.ReadFile(directory + "/publicKey")
	if err != nil {
		return "", "", err
	}
	return string(publicKey), string(privateKey), nil
}

func SaveSignatureToFile(signature string, filename string) error {
	return os.WriteFile(filename, []byte(signature), 0666)
}

func GetSignature(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
