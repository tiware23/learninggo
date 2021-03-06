package heighconv

import "fmt"

type Libras float64
type Kg float64

func (l Libras) String() string {
	return fmt.Sprintf("%gL", l)
}

func (k Kg) String() string {
	return fmt.Sprintf("%gkg", k)
}
