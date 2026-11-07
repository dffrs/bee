package internal

import (
	"errors"
	"fmt"
	"math"
	"strings"

	ut "bee/internal/util"
)

const window int = 10

func findInLine(haystack, needle string) ([]string, error) {
	idxFind := strings.Index(haystack, needle)
	if idxFind == -1 {
		return nil, errors.New("no matches found")
	}

	leftBound := int(math.Max(0, float64(idxFind-window)))
	rightStart := idxFind + len(needle)
	rightBound := int(math.Min(float64(len(haystack)), float64(rightStart+window)))

	left := haystack[leftBound:idxFind]
	right := haystack[rightStart:rightBound]

	needle = fmt.Sprintf(ut.Yellow+"%s"+ut.Reset, needle)

	return []string{left, needle, right}, nil
}

func findIn(haystack []string, needle string) (string, error) {
	needles := ""

	for _, hay := range haystack {
		match, err := findInLine(hay, needle)
		if err != nil {
			continue
		}

		needles += strings.Join(match, "") + "\n"
	}

	if len(needles) == 0 {
		return "", errors.New("no matches found")
	}

	return needles, nil
}
