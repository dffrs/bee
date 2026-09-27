package internal

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func findBeeDir() (string, error) {
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
			return io.EOF
		}

		return nil
	})
	if err != nil && err != io.EOF {
		return "", err
	}

	if beePath == "" {
		return "", errors.New("not found")
	}

	return beePath, nil
}

func tryGetBeeDir() (string, error) {
	var path string
	var err error

	result := make(chan string, 1)
	errCh := make(chan error, 1)

	go func() {
		path, err := findBeeDir()
		if err != nil {
			errCh <- err
			return
		}

		result <- path
	}()

	select {
	case path = <-result:
		return path, nil

	case err = <-errCh:
		return "", err

	case <-time.After(time.Second * 3):
		return "", errors.New("timed out trying to get bee directory")
	}
}

func buildContent(args []string) (string, error) {
	content := strings.Builder{}

	content.WriteString("[Summary]")
	content.WriteString("\n")

	for i, word := range args[1:] {
		var str string
		if i == len(args[1:])-1 {
			str = word
		} else {
			str = word + " "
		}

		if _, err := content.WriteString(str); err != nil {
			return "", errors.New("failed to write to string builder")
		}
	}

	content.WriteString("\n")

	return content.String(), nil
}

func createFile(dir string) (*os.File, string, error) {
	const createRetries = 5

	for range createRetries {
		fullPath := filepath.Join(dir, strconv.FormatInt(time.Now().UnixNano(), 10)+".bee")

		f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0o644)
		if errors.Is(err, fs.ErrExist) {
			continue
		}

		if err != nil {
			return nil, "", fmt.Errorf("failed to create new bee file (%s): %w", fullPath, err)
		}

		return f, fullPath, nil
	}

	return nil, "", fmt.Errorf("failed to create a new bee file in (%s) after %d attempts", dir, createRetries)
}

func (cm Command) New(args []string) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("summary is needed")
	}

	path, err := tryGetBeeDir()
	if err != nil {
		return "", err
	}

	content, err := buildContent(args)
	if err != nil {
		return "", err
	}

	newFile, newFilePath, err := createFile(path)
	if err != nil {
		return "", err
	}
	defer newFile.Close()

	_, err = newFile.WriteString(content)
	if err != nil {
		return "", fmt.Errorf("failed to write to file (%s): %w", newFilePath, err)
	}

	return fmt.Sprintf("created a new task (%s)\n", newFilePath), nil
}
