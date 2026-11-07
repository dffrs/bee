package internal

import (
	"errors"
	"fmt"
	"math"
	"strconv"
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

	leftBound := int(math.Max(0, float64(idxFind-window)))
	rightStart := idxFind + len(needle)
	rightBound := int(math.Min(float64(len(haystack)), float64(rightStart+window)))

	left := haystack[leftBound:idxFind]
	right := haystack[rightStart:rightBound]

	needle = fmt.Sprintf(ut.Yellow+"%s"+ut.Reset, needle)

	return []string{left, needle, right}, nil
}

func findIn(haystack []string, needle string) string {
	needles := ""

	for _, hay := range haystack {
		match, err := findInLine(hay, needle)
		if err != nil {
			continue
		}

		needles += strings.Join(match, "") + "\n"
	}

	return needles
}

func (cm Command) Find(args []string) (string, error) {
	if len(args) <= 1 {
		return "", errors.New("option is needed")
	}

	if len(args) <= 2 {
		return "", errors.New("value is needed")
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
		summaries := make([]string, 0)
		for _, duty := range duties {
			summaries = append(summaries, duty.Summary)
		}

		return findIn(summaries, toFind), nil
	case "priority":
		priorities := make([]string, 0)
		for _, duty := range duties {
			priorities = append(priorities, strconv.Itoa(duty.Priority))
		}

		return findIn(priorities, toFind), nil
	case "status":
		statuses := make([]string, 0)
		for _, duty := range duties {
			statuses = append(statuses, duty.Status)
		}

		return findIn(statuses, toFind), nil
	case "body":
		bodies := make([]string, 0)
		for _, duty := range duties {
			bodies = append(bodies, strings.Split(duty.Body, "\n")...)
		}

		return findIn(bodies, toFind), nil
	// case "any":

	default:
		return "", errors.New("option not recognized")
	}
}
