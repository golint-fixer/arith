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
	min := 1
	min <<= (unsafe.Sizeof(min)*8 - 1)
	max := min - 1

	tests := []struct {
		a, b int
		want bool
	}{
		{1, 2, true},
		{0, 0, true},
		{min, 0, true},
		{min, 1, true},
		{max, -1, true},
		{max, 0, true},
		{min, -1, false},
		{max, 1, false},
		{min, max, true},
	}

	for _, test := range tests {
		if got := IntAddOK(test.a, test.b); got != test.want {
			t.Errorf("IntAddOK(%d, %d) = %t\n", test.a, test.b, got)
		}
	}
}
