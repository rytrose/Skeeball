package util

import "runtime"

// IsRasPi returns true if the runtime is expected to be on a raspberry pi.
func IsRasPi() bool {
	// HACK: assume any linux runtime is a raspberry pi
	return runtime.GOOS == "linux"
}
