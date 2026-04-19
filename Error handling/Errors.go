package main

import(
	"fmt"
	"log"
)

func add(a int, b int) (int, error){
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("Negative numbers are not allowed: %d, %d", a, b)
	}
	return a + b, nil
}

func main(){
	result, err := add(5, -3)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Result:", result)
	}
}