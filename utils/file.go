package utils

import (
	"os"
)

func WriteToFile(filename string, content []byte) (error, string) {
	err := os.WriteFile(filename, content, 0600)
	if err != nil {
		return err, ""
	}
	return nil, filename
}
