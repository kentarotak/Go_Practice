package main

import "testing"

func TestShacountAllSame(t *testing.T) {

	var testval1, testval2 [32]byte
	expect := 0

	for i := 0; i < len(testval1); i++ {
		testval1[i] = 0xFF
		testval2[i] = 0xFF
	}

	if got := compareSHA(testval1, testval2); int(got) != expect {
		t.Errorf("err %v = %v", expect, got)
	}

}

func TestShacountAllCount(t *testing.T) {

	var testval1, testval2 [32]byte
	expect := 256

	for i := 0; i < len(testval1); i++ {
		testval1[i] = 0xFF
		testval2[i] = 0x00
	}

	if got := compareSHA(testval1, testval2); int(got) != expect {
		t.Errorf("err %v = %v", expect, got)
	}

}

func TestShacount(t *testing.T) {

	var testval1, testval2 [32]byte
	expect := 2

	for i := 0; i < len(testval1)-1; i++ {
		testval1[i] = 0xFF
		testval2[i] = 0xFF
	}

	testval1[31] = 0xEE
	testval2[31] = 0xCC

	if got := compareSHA(testval1, testval2); int(got) != expect {
		t.Errorf("err %v = %v", expect, got)
	}

}
