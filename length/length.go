package length

import (
	"github.com/gonum/unit"
)

// Represents a length in meters
type Length float64

const (
	Centimeter Length = 0.01
	Meter      Length = 1.0
	Foot       Length = 0.3048
)

func (l Length) OfUnit() *unit.Unit {
	return unit.Length(1)
}

func (l Length) Value() float64 {
	return float64(l)
}

func (l Length) Meters() float64 {
	return float64(l)
}

func (l Length) Feet() float64 {
	return float64(l) / float64(Foot)
}

func (l Length) Centimeters() float64 {
	return float64(l) / float64(Centimeter)
}
