package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Confirm(question string, defaultYes bool) (bool, error) {
	fmt.Println(question)

	// When stdin is not a terminal
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return confirmFromLine(defaultYes)
	}

	return confirmFromKey(defaultYes)
}

func confirmFromKey(defaultYes bool) (bool, error) {
	fd := int(os.Stdin.Fd())

	raw, err := term.MakeRaw(fd)
	if err != nil {
		return false, fmt.Errorf("failed to set terminal in raw mode: %w", err)
	}
	defer func() {
		// the terminal must be restored even if reading the answer fails,
		// otherwise the shell it runs in stays unusable.
		if err := term.Restore(fd, raw); err != nil {
			fmt.Fprintf(os.Stderr, "failed to restore terminal: %s\n", err)
		}
	}()

	for {
		key := make([]byte, 1)
		if _, err := os.Stdin.Read(key); err != nil {
			return false, fmt.Errorf("failed to read input: %w", err)
		}

		// the terminal does not echo the key back, so the newline the user did
		// not press is printed here to leave the cursor on a fresh line.
		switch key[0] {
		case 'y', 'Y':
			fmt.Println()
			return true, nil
		case 'n', 'N':
			fmt.Println()
			return false, nil
		case '\r', '\n':
			fmt.Println()
			return defaultYes, nil
		}

		fmt.Println("\nplease answer 'y' or 'n'")
	}
}

func confirmFromLine(defaultYes bool) (bool, error) {
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && line == "" {
		return false, fmt.Errorf("failed to read input: %w", err)
	}

	switch strings.ToLower(strings.TrimSpace(line)) {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	case "":
		return defaultYes, nil
	}

	return false, nil
}
