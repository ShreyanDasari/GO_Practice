package main

import(
	"fmt"
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
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}