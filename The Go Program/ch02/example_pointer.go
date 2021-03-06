// Examples of pointer.
package main

import "fmt"

func main() {
	var p = f()
        fmt.Println(p == p)
        fmt.Println(f() == f())
        b := new(int) // b, do tipo *int, aponta para uma variavel int sem nome
        fmt.Println(*b) // "0"
        *b = 2 // define o int sem nome com 2
        fmt.Println(*b)
}

func f() *int {
	v := 1
	return &v
}
