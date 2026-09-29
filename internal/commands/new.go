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

	dt "bee/internal/duty"
)

const (
	beeDir = ".bee"
	beeExt = ".md"
)

// TODO: move me someplace else
func findBeeDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	for {
		beePath := filepath.Join(dir, beeDir)

		info, err := os.Stat(beePath)
		if err == nil {
			if info.IsDir() {
				return beePath, nil
			}

			return "", fmt.Errorf("%s exists but is not a directory", beePath)
		}

		if !errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("check %s: %w", beePath, err)
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			break
		}

		dir = parent
	}

	return "", errors.New(".bee directory not found")
}

// TODO: move me someplace else
func tryGetBeeDir() (string, error) {
	result := make(chan struct {
		path string
		err  error
	}, 1)

	go func() {
		path, err := findBeeDir()

		result <- struct {
			path string
			err  error
		}{path: path, err: err}
	}()

	select {
	case res := <-result:
		return res.path, res.err

	case <-time.After(time.Second * 3):
		return "", errors.New("timed out trying to get bee directory")
	}
}

func buildContent(args []string) string {
	task := dt.NewDuty(strings.Join(args[1:], " "), 100, "TODO")

	return task.ToMarkdown()
}

func createFile(dir string) (*os.File, string, error) {
	const createRetries = 5

	for range createRetries {
		fullPath := filepath.Join(dir, strconv.FormatInt(time.Now().UnixNano(), 10)+beeExt)

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

	content := buildContent(args)

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
