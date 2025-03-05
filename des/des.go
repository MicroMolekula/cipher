package des

import (
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"
)

func Code(word, key string) (string, error) {
	binWord := ToBinary(word)
	binKey := ToBinary(key)
	binWords := make([][]int, 0)

	for w := range slices.Chunk(binWord, 64) {
		binWords = append(binWords, w)
	}
	if len(binWords[len(binWords)-1]) < 64 {
		binWords[len(binWords)-1] = append(make([]int, 64-len(binWords[len(binWords)-1])), binWords[len(binWords)-1]...)
	}
	cips := make([][]int, len(binWords))
	for i, w := range binWords {
		cips[i] = CodeBlock(w, binKey)
	}
	bits := ""
	for _, c := range cips {
		bits += Format(c)
	}
	result, err := ToHex(bits)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Decode(cip, key string) (string, error) {
	cips, err := FromHexToBinary(cip)
	if err != nil {
		return "", err
	}
	binKey := ToBinary(key)
	binWords := make([][]int, len(cips))
	for i, c := range cips {
		binWords[i] = DecodeBlock(c, binKey)
	}
	result := ""
	for _, w := range binWords {
		result += ToWord(w)
	}
	return result, nil
}

func CodeBlock(binWord, binKey []int) []int {
	ipWord := BlockPermutation(binWord, IP)
	ipKey := BlockPermutation(binKey, PC1)
	keyRound := ConvertKeyInRound(ipKey, 1)
	L, R := ArrToHalf(ipWord)
	for i := 0; i < 16; i++ {
		Ri := BlockPermutation(R, E)
		if i != 0 {
			keyRound = ConvertKeyInRound(keyRound, i+1)
		}
		afterPC2 := BlockPermutation(keyRound, PC2)
		xorKR := Xor(afterPC2, Ri)
		fromConvertS := FuncS(xorKR)
		fromF := BlockPermutation(fromConvertS, P)
		xorLR := Xor(L, fromF)
		L = R
		R = xorLR
	}
	return BlockPermutation(ArrayMerge(L, R), FP)
}

func DecodeBlock(cipWord, binKey []int) []int {
	keys := make([][]int, 16)
	ipKey := BlockPermutation(binKey, PC1)
	keys[0] = ConvertKeyInRound(ipKey, 1)
	for i := 1; i < 16; i++ {
		keys[i] = ConvertKeyInRound(keys[i-1], i+1)
	}
	ipCipWord := BlockPermutation(cipWord, IP)
	L, R := ArrToHalf(ipCipWord)
	for i := 15; i >= 0; i-- {
		Li := BlockPermutation(L, E)
		afterPC2 := BlockPermutation(keys[i], PC2)
		xorKL := Xor(Li, afterPC2)
		fromConvertS := FuncS(xorKL)
		fromP := BlockPermutation(fromConvertS, P)
		xorLR := Xor(R, fromP)
		R = L
		L = xorLR
	}
	return BlockPermutation(ArrayMerge(L, R), FP)
}

func ToWord(cip []int) string {
	chars := make([][]int, 0)
	for c := range slices.Chunk(cip, 8) {
		chars = append(chars, c)
	}
	word := ""
	for _, char := range chars {
		ch, _ := binaryToByte(char)
		if ch != 0 {
			word += string(byte(ch))
		}
	}
	return word
}

func ToHex(cip string) (string, error) {
	bigInt := big.NewInt(0)
	bigInt, ok := bigInt.SetString(cip, 2)
	if !ok {
		return "", errors.New("ошибка привода числа в строку")
	}
	return bigInt.Text(16), nil
}

func FromHexToBinary(hex string) ([][]int, error) {
	bigInt := big.NewInt(0)
	bigInt, ok := bigInt.SetString(hex, 16)
	if !ok {
		return nil, errors.New("ошибка привода строки в число")
	}
	bits := bigInt.Text(2)
	if len(bits)%64 != 0 {
		dopZero := ""
		for range 64 - len(bits)%64 {
			dopZero += "0"
		}
		bits = dopZero + bits
	}
	cipString := strings.Split(bits, "")
	cip := make([]int, len(cipString))
	for i, c := range cipString {
		ci, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		cip[i] = ci
	}
	result := make([][]int, 0)
	for c := range slices.Chunk(cip, 64) {
		result = append(result, c)
	}
	return result, nil
}

func ConvertKeyInRound(cipKey []int, numRound int) []int {
	arr1, arr2 := ArrToHalf(cipKey)
	afterPC1 := ArrayMerge(rotateLeft(arr1, KR[numRound-1]), rotateLeft(arr2, KR[numRound-1]))
	return afterPC1
}

func BlockPermutation(bits []int, tablePermutation []int) []int {
	result := make([]int, len(tablePermutation))
	for i, b := range tablePermutation {
		result[i] = bits[b]
	}
	return result
}

func ToBinary(word string) []int {
	wordBytes := []byte(word)
	resultString := make([]string, len(wordBytes))
	for i := 0; i < len(wordBytes); i++ {
		resultString[i] = fmt.Sprintf("%08b", wordBytes[i])
	}
	resultInt := make([]int, len(resultString)*8)
	for i := 0; i < len(resultString); i++ {
		for j := 0; j < len(resultString[i]); j++ {
			resultInt[i*8+j], _ = strconv.Atoi(string(resultString[i][j]))
		}
	}
	if len(resultInt) < 64 {
		dopInt := make([]int, 64-len(resultInt))
		resultInt = append(dopInt, resultInt...)
	}
	return resultInt
}

func ArrToHalf(array []int) ([]int, []int) {
	size := len(array) / 2
	arr1 := make([]int, size)
	arr2 := make([]int, size)
	for i := 0; i < len(array); i++ {
		switch {
		case i < size:
			arr1[i] = array[i]
		case i >= size:
			arr2[i-size] = array[i]
		}
	}
	return arr1, arr2
}

func rotateLeft(arr []int, shift int) []int {
	length := len(arr)
	shift = shift % length
	return append(arr[shift:], arr[:shift]...)
}

func LeftShift(array []int, value int) []int {
	b, _ := binaryToByte(array)
	bits := strings.Split(fmt.Sprintf("%08b", (b<<value)|(b>>(28-value))), "")
	if len(bits) > 28 {
		bits = strings.Split(fmt.Sprintf("%08b", ((b<<value)|(b>>(28-value)))&0x0FFFFFFF), "")
	}
	bitsShiftArr := make([]int, len(bits))
	for i := 0; i < len(bits); i++ {
		bitsShiftArr[i], _ = strconv.Atoi(bits[i])
	}
	if len(bitsShiftArr) < len(array) {
		dopInt := make([]int, len(array)-len(bitsShiftArr))
		bitsShiftArr = append(dopInt, bitsShiftArr...)
	}
	return bitsShiftArr
}

func ArrayMerge(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}

func Xor(arr1 []int, arr2 []int) []int {
	int1, err := binaryToByte(arr1)
	if err != nil {
		fmt.Println(err)
	}
	int2, err := binaryToByte(arr2)
	if err != nil {
		fmt.Println(err)
	}
	resultInt := int1 ^ int2
	result := stringToArrayInts(fmt.Sprintf("%08b", resultInt))
	if len(result) < len(arr1) {
		dopInt := make([]int, len(arr1)-len(result))
		result = append(dopInt, result...)
	}
	return result
}

func binaryToByte(bits []int) (int64, error) {
	binaryStr := ""
	for _, bit := range bits {
		binaryStr += fmt.Sprintf("%d", bit)
	}
	result, err := strconv.ParseInt(binaryStr, 2, len(bits)+1)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func stringToArrayInts(val string) []int {
	bits := strings.Split(val, "")
	result := make([]int, len(bits))
	for i := 0; i < len(bits); i++ {
		result[i], _ = strconv.Atoi(bits[i])
	}
	return result
}

func FuncS(val []int) []int {
	b := make([][]int, 0)
	for c := range slices.Chunk(val, 6) {
		b = append(b, c)
	}
	result := ""
	for i := 0; i < 8; i++ {
		result += ConvertS(b[i], S[i])
	}
	return stringToArrayInts(result)
}

func ConvertS(val []int, s [4][16]int) string {
	row := []int{val[0], val[5]}
	col := []int{val[1], val[2], val[3], val[4]}
	rowInt, _ := binaryToByte(row)
	colInt, _ := binaryToByte(col)
	sInt := s[rowInt][colInt]
	return fmt.Sprintf("%04b", sInt)
}

func Format(in []int) string {
	arr := make([]string, len(in))
	for i := 0; i < len(in); i++ {
		arr[i] = strconv.Itoa(in[i])
	}
	return strings.Join(arr, "")
}
