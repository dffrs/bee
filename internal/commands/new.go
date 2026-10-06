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
	ut "bee/internal/util"
)

func buildContent(args []string) string {
	task := dt.NewDuty(strings.Join(args[1:], " "), 100, "In Progress", "TODO")

	return task.ToMarkdown()
}

func createFile(dir string) (*os.File, string, error) {
	const createRetries = 5

	for range createRetries {
		fullPath := filepath.Join(dir, strconv.FormatInt(time.Now().UnixNano(), 10)+ut.BeeExt)

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

	path, err := ut.TryGetBeeDir()
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
