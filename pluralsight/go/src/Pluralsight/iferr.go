package main

import (
	"fmt"
	"os"
)

func main() {
	// _, err := os.Open("text.txt")
	// if err != nil {
	// 	fmt.Println("Error File: ", err)
	// }
	// Statement; Bool
	if _, err := os.Open("text.txt"); err != nil {
		fmt.Println("Error File: ", err)
	}
}
