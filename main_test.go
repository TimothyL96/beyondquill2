package main

import (
	"testing"
)

// func TestPow2(t *testing.T) {
// 	want := newLargeNumber("4")
// 	a := newLargeNumber("2")
// 	get := pow(a, a)
//
// 	if want.String() != get.String() {
// 		t.Errorf("Error in TestPow. Want %s, Get %s", want.String(), get.String())
// 	}
// }
//
// func TestPow3(t *testing.T) {
// 	want := newLargeNumber("27")
// 	a := newLargeNumber("3")
// 	get := pow(a, a)
//
// 	if want.String() != get.String() {
// 		t.Errorf("Error in TestPow. Want %s, Get %s", want.String(), get.String())
// 	}
// }
//
// func TestPow5(t *testing.T) {
// 	want := newLargeNumber("3125")
// 	a := newLargeNumber("5")
// 	b := newLargeNumber("5")
// 	get := pow(b, a)
//
// 	if want.String() != get.String() {
// 		t.Errorf("Error in TestPow. Want %s, Get %s", want.String(), get.String())
// 	} else {
// 		t.Log("Get", get.String())
// 	}
// }
//
// func TestPow12(t *testing.T) {
// 	want := newLargeNumber("8916100448256")
// 	a := newLargeNumber("12")
// 	b := newLargeNumber("12")
// 	get := pow(b, a)
//
// 	if want.String() != get.String() {
// 		t.Errorf("Error in TestPow. Want %s, Get %s", want.String(), get.String())
// 	} else {
// 		t.Log("Get", get.String())
// 	}
// }

func TestPow20(t *testing.T) {
	want := newLargeNumber("104857600000000000000000000")
	a := newLargeNumber("20")
	b := newLargeNumber("20")
	get := pow(b, a)

	if want.String() != get.String() {
		t.Errorf("Error in TestPow. Want %s, Get %s", want.String(), get.String())
	} else {
		t.Log("Get", get.String())
	}
}

// func BenchmarkPow(b *testing.B) {
// 	x := newLargeNumber("5")
// 	for i := 0; i < b.N; i++ {
// 		pow(x, x)
// 	}
// }
//
// func TestLargeNumber_Add(t *testing.T) {
// 	ln := newLargeNumber("100")
// 	ln1 := newLargeNumber("50")
// 	ln.Add(ln1)
//
// 	lnAns := newLargeNumber("150")
// 	get := ln.String()
// 	want := lnAns.String()
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber. Want %s, Get %s", want, get)
// 	} else {
// 		t.Log("Passed. Got:", get)
// 	}
// }
//
// func TestLargeNumber_Add1(t *testing.T) {
// 	first := "999999999999999999999999999999999999999"
// 	second := "2"
// 	ln := newLargeNumber(first)
// 	ln1 := newLargeNumber(second)
// 	ln.Add(ln1)
//
// 	get := ln.String()
// 	want := "1000000000000000000000000000000000000001"
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber. Want %s, Get %s", want, get)
// 	} else {
// 		t.Logf("Passed. Want %s + %s. Got %s", first, second, get)
// 	}
// }
//
// func TestLargeNumber_Add2(t *testing.T) {
// 	first := "3125"
// 	second := "288"
// 	ln := newLargeNumber(first)
// 	ln1 := newLargeNumber(second)
// 	ln.Add(ln1)
//
// 	get := ln.String()
// 	want := "3413"
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber. Want %s, Get %s", want, get)
// 	} else {
// 		t.Logf("Passed. Want %s + %s. Got %s", first, second, get)
// 	}
// }
//
// func TestLargeNumber_Add3(t *testing.T) {
// 	first := "8916100448256"
// 	second := "295716741928"
//
// 	ln := newLargeNumber(first)
// 	ln1 := newLargeNumber(second)
// 	ln.Add(ln1)
//
// 	get := ln.String()
// 	want := "9211817190184"
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber. Want %s, Get %s", want, get)
// 	} else {
// 		t.Logf("Passed. Want %s + %s. Got %s", first, second, get)
// 	}
// }
//
// func TestLargeNumber_Multiply(t *testing.T) {
// 	// First number
// 	ln := newLargeNumber("1000000000000000000000000000000000000000000000000000000000000000000")
//
// 	// Second number
// 	ln1 := newLargeNumber("2")
//
// 	// Multiple them
// 	ln.Multiply(ln1)
//
// 	// Wanted answer
// 	lnAns := newLargeNumber("2000000000000000000000000000000000000000000000000000000000000000000")
//
// 	// Set get and want
// 	get := ln.String()
// 	want := lnAns.String()
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber multiple. Want %s, Get %s", want, get)
// 	} else {
// 		t.Log("Passed. Got:", get)
// 	}
// }
//
// func TestLargeNumber_Multiply1(t *testing.T) {
// 	// First number
// 	ln := newLargeNumber("5159780352")
//
// 	// Second number
// 	ln1 := newLargeNumber("12")
//
// 	// Multiple them
// 	ln.Multiply(ln1)
//
// 	// Wanted answer
// 	lnAns := newLargeNumber("61917364224")
//
// 	// Set get and want
// 	get := ln.String()
// 	want := lnAns.String()
//
// 	if want != get {
// 		t.Errorf("Error in TestLargeNumber multiple. Want %s, Get %s", want, get)
// 	} else {
// 		t.Log("Passed. Got:", get)
// 	}
// }
