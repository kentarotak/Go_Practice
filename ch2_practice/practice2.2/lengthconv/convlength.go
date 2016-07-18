
package lengthconv

type Feet float64
type Meter float64

func FtoM(f Feet) Meter { return Meter(f/3.2808)}

func MtoF(m Meter) Feet { return Feet(m*3.2808)}

//!-
