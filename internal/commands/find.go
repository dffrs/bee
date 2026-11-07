package internal

import (
	"errors"
	"strconv"
	"strings"

	d "bee/internal/duty"
	ut "bee/internal/util"
)

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

		return findIn(summaries, toFind)
	case "priority":
		priorities := make([]string, 0)
		for _, duty := range duties {
			priorities = append(priorities, strconv.Itoa(duty.Priority))
		}

		return findIn(priorities, toFind)
	case "status":
		statuses := make([]string, 0)
		for _, duty := range duties {
			statuses = append(statuses, duty.Status)
		}

		return findIn(statuses, toFind)
	case "body":
		bodies := make([]string, 0)
		for _, duty := range duties {
			bodies = append(bodies, strings.Split(duty.Body, "\n")...)
		}

		return findIn(bodies, toFind)
	// case "any":

	default:
		return "", errors.New("option not recognized")
	}
}
