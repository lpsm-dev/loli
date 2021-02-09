package helpers

import "os"

// FileExists function - check with a file exist in system.
func FileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
