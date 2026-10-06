package internal

import (
	"fmt"
	"os"

	ut "bee/internal/util"
)

func (cm Command) Purge(args []string) (string, error) {
	path, err := ut.TryGetBeeDir()
	if err != nil {
		return "", err
	}

	fmt.Printf(ut.Yellow+"[Bee]"+ut.Reset+"\nfound bee dir in: %s\n", path)

	confirmed, err := ut.Confirm("\ndo you want to delete it ? [Y/n]", true)
	if err != nil {
		return "", err
	}

	if !confirmed {
		return "aborting...\n", nil
	}

	if err := os.RemoveAll(path); err != nil {
		return "", fmt.Errorf("failed to delete bee dir: %w", err)
	}

	return "deleted bee dir\n", nil
}
