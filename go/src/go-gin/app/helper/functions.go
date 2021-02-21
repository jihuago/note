package helper

import (
	"os"
	"path/filepath"
)

func config()  {
	
}

func GetCurrentPath() string  {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}

	return "./"
}
