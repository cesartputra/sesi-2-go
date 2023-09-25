package main

import "fmt"

func main() {
	str := "selamat malam"

	charMap := make(map[string]int)

	for _, char := range str {
		fmt.Println(string(char))
		charMap[string(char)]++
	}

	fmt.Println(charMap)
}
