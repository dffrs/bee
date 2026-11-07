package internal

import (
	"errors"
	"fmt"
	"math"
	"strings"

	d "bee/internal/duty"
	ut "bee/internal/util"
)

const window int = 10

func findInLine(haystack, needle string) ([]string, error) {
	idxFind := strings.Index(haystack, needle)
	if idxFind == -1 {
		return nil, errors.New("no matches found")
	}

	leftBound := math.Max(0, float64(idxFind-window))
	rightBound := math.Min(float64(len(haystack)-1), float64(idxFind+len(needle)-1+window))

	temp := idxFind + len(needle)
	if temp >= len(haystack) {
		temp = len(haystack) - 1
	}

	left := haystack[int(leftBound):idxFind]
	right := haystack[temp:int(rightBound)]

	needle = fmt.Sprintf(ut.Yellow+"%s"+ut.Reset, needle)

	return []string{left, needle, right}, nil
}

func findInSummary(duties []d.Duty, toFind string) string {
	nail := ""

	for _, duty := range duties {
		match, err := findInLine(duty.Summary, toFind)
		if err != nil {
			continue
		}

		nail = strings.Join(match, "") + "\n"
	}

	return nail
}

func (cm Command) Find(args []string) (string, error) {
	if len(args) <= 1 {
		// TODO: get proper error messsage
		return "", errors.New("option is needed")
	}

	path, err := ut.TryGetBeeDir()
	if err != nil {
		return "", err
	}

	duties := make([]d.Duty, 0)
	if err := findMarkdownFiles(path, &duties); err != nil {
		return "", err
	}

	if len(duties) == 0 {
		return "not duties found. Run 'bee new <summary>' to create a new duty\n", nil
	}

	toFind := args[2]
	option := args[1]

	switch option {
	case "summary":
		return findInSummary(duties, toFind), nil
	// case "priority":
	// case "status":
	// case "body":
	// case "any":

	default:
		return "", errors.New("option not recognized")
	}
}
