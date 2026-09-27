package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

	content := strings.Builder{}

	content.WriteString("[Summary]")
	content.WriteString("\n")

	for _, word := range args[1:] {
		if _, err := content.WriteString(word + " "); err != nil {
			return "", errors.New("failed to write to string builder")
		}
	}

	content.WriteString("\n")

	timestamp := int(time.Now().Unix())
	newFilePath := filepath.Join(path, strconv.Itoa(timestamp))

	f, err := os.Create(fmt.Sprintf("%s.bee", newFilePath))
	if err != nil {
		return "", fmt.Errorf("failed to create new bee file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(content.String())
	if err != nil {
		return "", fmt.Errorf("failed to write to file (%s): %w", newFilePath, err)
	}

	return fmt.Sprintf("created a new task (%s)\n", newFilePath), nil
}
