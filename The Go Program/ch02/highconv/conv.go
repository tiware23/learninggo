package highconv

// MToF converts Meters to Feet
func MToF(m Meters) Feet {
	return Feet(m * 2)
}

// FToM converts Feet to Meters
func FToM(f Feet) Meters {
	return Meters(f / 3.280)
}
