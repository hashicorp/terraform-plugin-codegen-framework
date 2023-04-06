package input

import (
	"os"
)

func Read(path string) ([]byte, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return src, nil
}
