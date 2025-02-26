package config

import "os"

func fileExists(filename string) bool {
	_, err := os.Lstat(filename)
	return !os.IsNotExist(err)
}
