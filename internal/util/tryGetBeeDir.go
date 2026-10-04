package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

const (
	beeDir = ".bee"
)

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

func TryGetBeeDir() (string, error) {
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
