package arith

// UIntAddOK test whether two uint add is overflow.
func UIntAddOK(a, b uint) bool {
	if a+b < a {
		return false
	}
	return true
}

// IntAddOK test whether two uint add is overflow or underflow.
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
