package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	d "bee/internal/duty"
	ut "bee/internal/util"
)

// ////////////////////////////////////////
func findInSummary(duties []d.Duty, toFind string) (string, error) {
	summaries := make([]string, 0)
	for _, duty := range duties {
		summaries = append(summaries, duty.Summary)
	}

	return findIn(summaries, toFind)
}

func findInPriorities(duties []d.Duty, toFind string) (string, error) {
	priorities := make([]string, 0)
	for _, duty := range duties {
		priorities = append(priorities, strconv.Itoa(duty.Priority))
	}

	return findIn(priorities, toFind)
}

func findInStatus(duties []d.Duty, toFind string) (string, error) {
	statuses := make([]string, 0)
	for _, duty := range duties {
		statuses = append(statuses, duty.Status)
	}

	return findIn(statuses, toFind)
}

func findInBody(duties []d.Duty, toFind string) (string, error) {
	bodies := make([]string, 0)
	for _, duty := range duties {
		bodies = append(bodies, strings.Split(duty.Body, "\n")...)
	}

	return findIn(bodies, toFind)
}

func findInAny(duties []d.Duty, toFind string) (string, error) {
	needles := ""

	summNeedle, err := findInSummary(duties, toFind)
	if err == nil {
		needles += fmt.Sprintf(ut.Magenta+"%s"+ut.Reset+"\n%s", "Summary", summNeedle)
	}

	prioNeedle, err := findInPriorities(duties, toFind)
	if err == nil {
		needles += fmt.Sprintf(ut.Magenta+"%s"+ut.Reset+"\n%s", "Priority", prioNeedle)
	}

	statusNeedle, err := findInStatus(duties, toFind)
	if err == nil {
		needles += fmt.Sprintf(ut.Magenta+"%s"+ut.Reset+"\n%s", "Status", statusNeedle)
	}

	bodyNeedle, err := findInBody(duties, toFind)
	if err == nil {
		needles += fmt.Sprintf(ut.Magenta+"%s"+ut.Reset+"\n%s", "Body", bodyNeedle)
	}

	if len(needles) == 0 {
		return "", errors.New("match not found")
	}

	return needles, nil
}

//////////////////////////////////////////

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
		return findInSummary(duties, toFind)
	case "priority":
		return findInPriorities(duties, toFind)
	case "status":
		return findInStatus(duties, toFind)
	case "body":
		return findInBody(duties, toFind)
	case "any":
		return findInAny(duties, toFind)
	default:
		return "", errors.New("option not recognized")
	}
}
