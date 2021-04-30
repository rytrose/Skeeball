package util

// Mod provides round-toward-zero modulus functionality.
// From https://stackoverflow.com/a/43018347.
func Mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

// Min returns the lesser of two ints.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the greater of two ints.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
