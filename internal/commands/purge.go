package internal

import (
	"bufio"
	"fmt"
	"os"

	ut "bee/internal/util"
)

func (cm Command) Purge(args []string) (string, error) {
	path, err := ut.TryGetBeeDir()
	if err != nil {
		return "", err
	}

	fmt.Printf("[Bee]\nfound bee dir in: %s\n", path)
	fmt.Print("\ndo you want to delete it ? [Y/n]")

	reader := bufio.NewReader(os.Stdin)

	choise, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	if choise != "y" && choise != "Y" && choise != "\n" {
		return "aborting...", nil
	}

	if err := os.RemoveAll(path); err != nil {
		return "", fmt.Errorf("failed to delete bee dir: %w", err)
	}

	return "deleted bee dir\n", nil
}
