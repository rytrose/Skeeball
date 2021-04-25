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
