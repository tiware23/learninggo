package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1:]
	if len(command) < 1 {
		fmt.Println("Missed args")
	}
	for _, i := range command {
		switch i {
		case "linux":
			fmt.Println("Linux")
		case "docker":
			fmt.Println("docker")
		case "Windows":
			fmt.Println("Ruindows")
		default:
			fmt.Println("See the top 100 courses available!")
		}
	}
}
