package internal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func exists() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var beePath string

	err = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".bee" {
			beePath = path
			return nil
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	if beePath == "" {
		return "", errors.New("not found")
	}

	return beePath, nil
}

func (cm Command) New(args []string) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("summary is needed")
	}

	// TODO: Maybe gorountine with select ?
	path, err := exists()
	if err != nil {
		return "", errors.New("failed to get bee directory")
	}

	summary := strings.Builder{}

	for _, word := range args[1:] {
		if _, err := summary.WriteString(word + " "); err != nil {
			return "", errors.New("failed to write to string builder")
		}
	}

	// NOTE: path already includes '.bee' dir
	_ = path

	return "added new task", nil
}
