package unit

import (
	"fmt"
	"strings"
)

// A uniter is a type that can be converted
// to a unit. These functions declare the
// power of the basic dimension of the unit
type Uniter interface {
	OfUnit() *Unit
}

type Quantitier interface {
	Uniter
	Value() float64
}

type unitType int

const (
	length unitType = iota
	mass
	time
	current
	temperature
	_
	luminosity
)

func (u unitType) String() string {
	switch u {
	case length:
		return "Length"
	case mass:
		return "Mass"
	case time:
		return "Time"
	case current:
		return "Current"
	case temperature:
		return "Temperature"
	case luminosity:
		return "Luminosity"
	}
	return fmt.Sprint(u)
}

type dimensions map[unitType]int

type Unit struct {
	dims  dimensions
}

func (u *Unit) OfUnit() *Unit {
	return u
}
func (u Unit) String() (s string) {
	str := make([]string, 0)
	for ut, dim := range u.dims {
		str = append(str, ut.String(), fmt.Sprintf(":%v ", dim)) 
	}
	return strings.Join(str, "")
}

func createUnit(d dimensions) *Unit {
	return &Unit{dims: d}
}

func Length(dim int) *Unit {
	return createUnit(dimensions{length: dim})
}

func Mass(dim int) *Unit {
	return createUnit(dimensions{mass: dim})
}

func Time(dim int) *Unit {
	return createUnit(dimensions{time: dim})
}

func Current(dim int) *Unit {
	return createUnit(dimensions{current: dim})
}

func Temperature(dim int) *Unit {
	return createUnit(dimensions{temperature: dim})
}

func Luminosity(dim int) *Unit {
	return createUnit(dimensions{luminosity: dim})
}

func DimensionsMatch(aU, bU Uniter) bool {
	a := aU.OfUnit()
	b := bU.OfUnit()
	if len(a.dims) != len(b.dims) {
		return false
	}
	for ut, dim := range a.dims {
		if dim != b.dims[ut] {
			return false
		}
	}
	return true
}

func Mul(aU, bU Uniter) {
	a := aU.OfUnit()
	b := bU.OfUnit()
	for k, v := range b.dims {
		_, ok := a.dims[k]
		if ok {
			a.dims[k] = v
			if a.dims[k] == 0 { delete(a.dims, k) }
		} else {
			a.dims[k] = v
		}
	}
	return
}

func Div(aU, bU Uniter) {
	a := aU.OfUnit()
	b := bU.OfUnit()
	for k, v := range b.dims {
		_, ok := a.dims[k]
		if ok {
			a.dims[k] -= v
			if a.dims[k] == 0 { delete(a.dims, k) }
		} else {
			a.dims[k] = -v
		}
	}
	return
}

type Quantity struct {
	Unit
	value float64
}

func (q *Quantity) Value() float64 {
	return q.value
}

func CreateQuantity(v float64, u *Unit) (q *Quantity) {
	q = &Quantity{*u, v}
	return q
}
