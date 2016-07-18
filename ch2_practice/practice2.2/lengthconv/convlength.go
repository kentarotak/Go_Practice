package lengthconv

import "fmt"

type Feet float64
type Meter float64

func FtoM(f Feet) Meter { return Meter(f / 3.2808) }

func MtoF(m Meter) Feet { return Feet(m * 3.2808) }

func (f Feet) String() string  { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }

//!-
