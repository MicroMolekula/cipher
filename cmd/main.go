package main

import (
	"fmt"
	"github.com/cipher/des"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//word := des.ToBinary("zalupa")
	//key := des.ToBinary("key")
	//cipWord := des.BlockPermutation(word, des.IP)
	//cipKey := des.BlockPermutation(key, des.PC1)
	//arr1, arr2 := des.ArrToHalf(cipKey)
	//afterPC1 := des.ArrayMerge(des.LeftShift(arr1, 1), des.LeftShift(arr2, 1))
	//afterPC2 := des.BlockPermutation(afterPC1, des.PC2)
	//L, R := des.ArrToHalf(cipWord)
	//Re := des.BlockPermutation(R, des.E)
	//xorKR0 := des.Xor(afterPC2, Re)
	//fmt.Println(format(xorKR0))
	//fromConvertS := des.FuncS(xorKR0)
	//fmt.Println(format(fromConvertS))
	//fromF := des.BlockPermutation(fromConvertS, des.P)
	////fmt.Println(format(fromF))
	//fmt.Println(format(des.Xor(L, fromF)))
	code, binWord := des.Code("zalupa", "key")
	decode := des.Decode(code, "key")
	if Format(binWord) == Format(decode) {
		fmt.Println("работает заебись")
	}
	// 536870912
}

func reverse() {
	ip := des.FP
	slices.Reverse[[]int, int](ip)
	result := make([]string, len(ip))
	for i := 0; i < len(ip); i++ {
		result[i] = strconv.Itoa(ip[i])
	}
	fmt.Println(strings.Join(result, ", "))
}

func Format(in []int) string {
	arr := make([]string, len(in))
	for i := 0; i < len(in); i++ {
		arr[i] = strconv.Itoa(in[i])
	}
	return strings.Join(arr, "")
}
