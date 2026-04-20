package main

import (
	"fmt"
)

func main(){
	str1 := "Hello, "
	str2 := "World!"
	result := str1 + str2
	runes := []rune(result)
	fmt.Println(result)
	for i, char := range result {
		fmt.Printf("Character %d: %c\n", i, char)
	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("Byte %d: %c\n", i, result[i])
	}
	for i, char := range runes {
		fmt.Printf("Rune %d: %c\n", i, char)
	}
}
