// ex2.2 read from args and stdin and convert from C to F, from feet to meters
// from peso to libras and kg.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

type Meters float64
type Feet float64

type Libras float64
type Kg float64

func MToF(m Meters) Feet {
	return Feet(m * 2)
}

func FToM(f Feet) Meters {
	return Meters(f / 3.280)
}

func LToK(l Libras) Kg {
	return Kg(l / 0.4536)
}

func KToL(k Kg) Libras {
	return Libras(k * 0.4536)
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
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			m := Meters(t)
			F := Feet(t)
			l := Libras(t)
			k := Kg(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
			fmt.Printf("%0.2f = %0.2f, %0.2f = %0.2f\n", m, MToF(m), F, FToM(F))
			fmt.Printf("%0.2f = %0.2f, %0.2f = %0.2f\n", l, LToK(l), k, KToL(k))
		}
	} else {
		for _, arg := range numbers {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			m := Meters(t)
			F := Feet(t)
			l := Libras(t)
			k := Kg(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
			fmt.Printf("%0.2f = %0.2f, %0.2f = %0.2f\n", m, MToF(m), F, FToM(F))
			fmt.Printf("%0.2f = %0.2f, %0.2f = %0.2f\n", l, LToK(l), k, KToL(k))
		}
	}
}
