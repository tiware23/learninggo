package heighconv

// LToK converts Libras to Kg
func LToK(l Libras) Kg {
	return Kg(l / 0.4536)
}

// KToL converts Kg to Libras
func KToL(k Kg) Libras {
	return Libras(k * 0.4536)
}
