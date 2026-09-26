package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

type Start struct {
	Command
}

func New() *Start {
	s := Start{}
	s.Name = "start"

	return &Start{}
}

func (st Start) Run(args []string) (string, error) {
	path := args[0]

	if path == "" {
		path = "./"
	}

	dirName := filepath.Join(path, "bee")

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("start bee dir in: %s", dirName), nil
}
