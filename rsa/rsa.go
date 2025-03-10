package rsa

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Key struct {
	Kn *big.Int
	N  *big.Int
}

func Encrypt(word string, key string) (string, error) {
	enKey, err := HexToKey(key)
	if err != nil {
		return "", err
	}
	arrayInt := WordToArrayInt(word)
	result := make([]string, len(arrayInt))
	for i, item := range arrayInt {
		result[i] = IntToHex(Code(item, enKey))
	}
	return strings.Join(result, "/"), nil
}

func Decrypt(word string, key string) (string, error) {
	deKey, err := HexToKey(key)
	if err != nil {
		return "", err
	}
	arrayInt, err := HexStringToArrayInt(word)
	if err != nil {
		return "", err
	}
	resultDecode := make([]*big.Int, len(arrayInt))
	for i, item := range arrayInt {
		resultDecode[i] = Decode(item, deKey)
	}
	result := make([]byte, len(resultDecode))
	for i, item := range resultDecode {
		result[i] = byte(item.Int64())
	}
	return string(result), nil
}

func Test() {
	fmt.Println(searchD(big.NewInt(17), big.NewInt(2651394840)).Text(10))
}

func Code(num *big.Int, key *Key) *big.Int {
	cip := num.Exp(num, key.Kn, nil)
	cip.Mod(cip, key.N)
	return cip
}

func Decode(number *big.Int, key *Key) *big.Int {
	return number.Exp(number, key.Kn, key.N)
}

func GenerateKeys(bits int) (string, string, error) {
	p, err := generatePrime(bits)
	if err != nil {
		return "", "", err
	}
	q, err := generatePrime(bits)
	if err != nil {
		return "", "", err
	}
	if p.Text(10) == q.Text(10) {
		p, err = generatePrime(bits)
		if err != nil {
			return "", "", err
		}
	}
	n := big.NewInt(1)
	n = n.Mul(p, q)
	subP := big.NewInt(1).Sub(p, big.NewInt(1))
	subQ := big.NewInt(1).Sub(q, big.NewInt(1))
	fn := big.NewInt(1).Mul(subP, subQ)
	e := searchE(fn)
	if e == nil {
		return "", "", errors.New("error generate e")
	}
	d := searchD(e, fn)
	if d == nil {
		return "", "", errors.New("error generate d")
	}
	keyEnc := &Key{
		Kn: e,
		N:  n,
	}
	keyDec := &Key{
		Kn: d,
		N:  n,
	}
	return KeyToHex(keyEnc), KeyToHex(keyDec), nil
}

func generatePrime(bits int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return prime, nil
}

func searchE(fn *big.Int) *big.Int {
	for i := 0; ; i++ {
		exp := big.NewInt(2)
		exp = exp.Exp(exp, big.NewInt(int64(i)), nil)
		e := big.NewInt(2)
		e = e.Exp(e, exp, nil)
		eR := *e.Add(e, big.NewInt(1))
		if e.GCD(nil, nil, e, fn).Text(10) == big.NewInt(1).Text(10) {
			return &eR
		}
		if e.Cmp(fn) == 1 {
			break
		}
	}
	return nil
}

func searchD(e, fn *big.Int) *big.Int {
	return big.NewInt(1).ModInverse(e, fn)
}

func IntToHex(num *big.Int) string {
	return num.Text(16)
}

func HexToInt(hex string) (*big.Int, error) {
	result, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		return nil, errors.New("error convert hex to int")
	}
	return result, nil
}

func ByteToInt(char byte) *big.Int {
	return big.NewInt(int64(char))
}

func WordToArrayInt(word string) []*big.Int {
	wordChars := []byte(word)
	result := make([]*big.Int, len(wordChars))
	for i, char := range wordChars {
		result[i] = ByteToInt(char)
	}
	return result
}

func HexStringToArrayInt(cip string) ([]*big.Int, error) {
	hexArray := strings.Split(cip, "/")
	result := make([]*big.Int, len(hexArray))
	for i, hex := range hexArray {
		bint, err := HexToInt(hex)
		if err != nil {
			return nil, err
		}
		result[i] = bint
	}
	return result, nil
}

func KeyToHex(key *Key) string {
	kn := IntToHex(key.Kn)
	n := IntToHex(key.N)
	return strings.Join([]string{kn, n}, ":")
}

func HexToKey(hex string) (*Key, error) {
	arr := strings.Split(hex, ":")
	if len(arr) != 2 {
		return nil, errors.New("error convert hex to key")
	}
	kn, err := HexToInt(arr[0])
	if err != nil {
		return nil, err
	}
	n, err := HexToInt(arr[1])
	if err != nil {
		return nil, err
	}
	return &Key{
		Kn: kn,
		N:  n,
	}, nil
}
