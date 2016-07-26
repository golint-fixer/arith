package arith

import (
	"testing"
	"unsafe"
)

var (
	base = 1
	tmin = base << (unsafe.Sizeof(base)*8 - 1)
	tmax = tmin - 1
)

func TestUIntAddOK(t *testing.T) {
	var dummy = ^0 // twos' complement -1
	var max = uint(dummy)
	tests := []struct {
		a, b uint
		want bool
	}{
		{1, 2, true},
		{max, 0, true},
		{max, 1, false},
		{max, 123, false},
	}

	for _, test := range tests {
		if got := UIntAddOK(test.a, test.b); got != test.want {
			t.Errorf("UAddOK(%d, %d) = %t\n", test.a, test.b, got)
		}
	}
}

func TestIntAddOK(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{1, 2, true},
		{0, 0, true},
		{tmin, 0, true},
		{tmin, 1, true},
		{tmax, -1, true},
		{tmax, 0, true},
		{tmin, -1, false},
		{tmax, 1, false},
		{tmin, tmax, true},
	}

	for _, test := range tests {
		if got := IntAddOK(test.a, test.b); got != test.want {
			t.Errorf("IntAddOK(%d, %d) = %t\n", test.a, test.b, got)
		}
	}
}

func TestIntSubOK(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{0, 0, true},
		{123, 0, true},
		{tmin, 0, true},
		{0, tmin, false},
		{1, tmin, false},
	}

	for _, test := range tests {
		if got := IntSubOK(test.a, test.b); got != test.want {
			t.Errorf("IntSubOK(%d, %d) = %t\n", test.a, test.b, got)
		}
	}
}

func TestIntMultOK(t *testing.T) {
	tests := []struct {
		a, b int
		want bool
	}{
		{0, 0, true},
		{0, 123, true},
		{tmax, 1, true},
		{tmax, 2, false},
		{tmin, 2, false},
		{tmin, tmax, false},
	}

	for _, test := range tests {
		if got := IntMultOK(test.a, test.b); got != test.want {
			t.Errorf("IntMultOK(%d, %d) = %t\n", test.a, test.b, got)
		}
	}
}

func TestDivByPow2(t *testing.T) {
	tests := []struct {
		a    int
		k    byte
		want int
	}{
		{12340, 0, 12340}, // 1
		{12340, 1, 6170},  // 2
		{12340, 4, 771},   // 16
		{-12340, 0, -12340},
		{-12340, 1, -6170},
		{-12340, 4, -771},
		{-12340, 8, -48}, // 256
	}

	for _, test := range tests {
		if got := DivByPow2(test.a, test.k); got != test.want {
			t.Errorf("DivByPow2(%d, %d) = %d, want %d\n", test.a, test.k, got, test.want)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		a    int
		want uint
	}{
		{0, 0},
		{1, 1},
		{tmax, uint(tmax)},
		{tmin, uint(tmin)},
		{-1, 1},
	}

	for _, test := range tests {
		if got := Abs(test.a); got != test.want {
			t.Errorf("Abs(%d) = %d, want %d\n", test.a, got, test.want)
		}
	}
}
