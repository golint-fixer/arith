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

// IntMultOK testing whether a*b is overflow.
func IntMultOK(a, b int) bool {
	if a == 0 {
		return true
	}
	p := a * b
	if p/a != b {
		return false
	}
	return true
}

// DivByPow2 fast a/2^k calculating.
func DivByPow2(a int, k byte) int {
	if a >= 0 {
		return a >> k
	}
	return (a + (1 << k) - 1) >> k
}

// Abs from int to uint
func Abs(a int) uint {
	if a > 0 || a == -a { // 0 or TMin
		return uint(a)
	}
	return uint(-a)
}
