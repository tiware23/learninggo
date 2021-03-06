package main

import (
	"fmt"
	"tempconv"
)

func main() {
	// fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC) // "Brrr! - 273.15°C"
	// fmt.Println(tempconv.CToF(tempconv.BoilingC))    // "212°F"
	// fmt.Println(tempconv.CToK(tempconv.BoilingC))
	fmt.Println(tempconv.FtoK(68))
}
