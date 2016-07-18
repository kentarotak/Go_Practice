package weightconv

type pound float64
type gram float64

func PtoG(p pound) gram { return gram(p * 453.592) }

func GtoP(g gram) pound { return pound(g / 453.592) }

//!-
