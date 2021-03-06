// ex2.2 read from args and stdin and convert from C to F, from feet to meters
// from peso to libras and kg.
package main

import (
	"bufio"
	"fmt"
	"heighconv"
	"highconv"
	"os"
	"strconv"
	"tempconv"
)

func ProcessArgs(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := highconv.Meters(t)
	F := highconv.Feet(t)
	l := heighconv.Libras(t)
	k := heighconv.Kg(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", m, highconv.MToF(m), F, highconv.FToM(F))
	fmt.Printf("%0.2f = %0.2f, %0.2f = %0.2f\n", l, heighconv.LToK(l), k, heighconv.KToL(k))
}

func main() {
	numbers := os.Args[1:]
	if len(numbers) < 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			ProcessArgs(t)
		}
	} else {
		for _, arg := range numbers {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			ProcessArgs(t)
		}
	}
}
