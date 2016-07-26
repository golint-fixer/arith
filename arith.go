package arith

// UIntAddOK test whether a + b is overflow.
func UIntAddOK(a, b uint) bool {
	if a+b < a {
		return false
	}
	return true
}

// IntAddOK test whether a + b is overflow or underflow.
func IntAddOK(a, b int) bool {
	sum := a + b
	if a >= 0 && b >= 0 && sum < 0 {
		return false
	}
	if a < 0 && b < 0 && sum >= 0 {
		return false
	}
	return true
}

// IntSubOK test whether a - b is overflow or underflow.
func IntSubOK(a, b int) bool {
	if b == 0 {
		return true
	}
	// any number - TMin is overflow
	if b == -b { // b == TMin
		return false
	}
	return IntAddOK(a, -b)
}
