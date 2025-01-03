package config

import (
	"errors"
	"os"
)

func CheckFileExists(path string) bool {
	_, error := os.Stat(path)

	return !errors.Is(error, os.ErrNotExist)
}
