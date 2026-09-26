package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func (st Command) Start(args []string) (string, error) {
	var path string

	if len(args) <= 1 {
		path = "./"
	} else {
		path = args[1]
	}

	dirName := filepath.Join(path, "bee")

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("bee active in: %s\n", dirName), nil
}
