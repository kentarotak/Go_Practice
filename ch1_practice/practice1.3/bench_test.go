package main

import "testing"

var testsize = 10000

func BenchmarkPoorFunc(b *testing.B) {

	args := []string{}

	for j := 0; j < testsize; j++ {
		args = append(args, string(j))
	}

	for i := 0; i < b.N; i++ {
		poorLogic(args)
	}
}

func BenchmarkNormalFunc(b *testing.B) {

	args := []string{}

	for j := 0; j < testsize; j++ {
		args = append(args, string(j))
	}

	for i := 0; i < b.N; i++ {
		normalLogic(args)
	}
}
