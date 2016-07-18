package weightconv

import "fmt"

type Pound float64
type Gram float64

func PtoG(p Pound) Gram { return Gram(p * 453.592) }

func GtoP(g Gram) Pound { return Pound(g / 453.592) }

func (p Pound) String() string { return fmt.Sprintf("%g pound", p) }
func (g Gram) String() string  { return fmt.Sprintf("%g gram", g) }

//!-
