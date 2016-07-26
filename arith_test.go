package arith

import (
	"testing"
	"unsafe"
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
			t.Errorf("UAddOK(%u, %u) = %t\n", test.a, test.b, got)
		}
	}
}

func TestIntAddOK(t *testing.T) {
	tmin := 1
	tmin <<= (unsafe.Sizeof(tmin)*8 - 1)
	tmax := tmin - 1

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
	tmin := 1
	tmin <<= (unsafe.Sizeof(tmin)*8 - 1)

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
