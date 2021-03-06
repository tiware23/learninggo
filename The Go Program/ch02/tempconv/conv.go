package tempconv

// CtoF converte uma temperatura em Celsius para Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converte uma temperatura em Fahrenheit para Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converte uma temperatura em Celsius para Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273)
}

func FtoK(f Fahrenheit) Kelvin {
	fToC := Celsius((f - 32) * 5 / 9)
	return Kelvin(fToC + 273)
}
