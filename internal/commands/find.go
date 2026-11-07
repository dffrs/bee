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
func findInDuty(duty d.Duty, option string, toFind string) (string, error) {
	switch option {
	case "summary":
		return findIn([]string{duty.Summary}, toFind)
	case "priority":
		return findIn([]string{strconv.Itoa(duty.Priority)}, toFind)
	case "status":
		return findIn([]string{duty.Status}, toFind)
	case "body":
		return findIn(strings.Split(duty.Body, "\n"), toFind)
	case "any":
		return findInAnyDuty(duty, toFind)
	default:
		return "", errors.New("option not recognized")
	}
}

func findInAnyDuty(duty d.Duty, toFind string) (string, error) {
	fields := []struct {
		name   string
		values []string
	}{
		{"Summary", []string{duty.Summary}},
		{"Priority", []string{strconv.Itoa(duty.Priority)}},
		{"Status", []string{duty.Status}},
		{"Body", strings.Split(duty.Body, "\n")},
	}

	needles := ""
	for _, field := range fields {
		matches, err := findIn(field.values, toFind)
		if err == nil {
			needles += fmt.Sprintf(ut.Magenta+"%s"+ut.Reset+"\n%s", field.name, matches)
		}
	}

	if len(needles) == 0 {
		return "", errors.New("match not found")
	}

	return needles, nil
}

func findInDuties(duties []d.Duty, option string, toFind string) (string, error) {
	needles := ""

	for _, duty := range duties {
		matches, err := findInDuty(duty, option, toFind)
		if err != nil {
			continue
		}

		if len(needles) > 0 {
			needles += "\n--------------------------------\n\n"
		}

		needles += matches
		needles += fmt.Sprintf(ut.Gray+"\nFile\n%s\n"+ut.Reset, duty.Path)
	}

	if len(needles) == 0 {
		return "", errors.New("match not found")
	}

	return needles, nil
}

func findInSummary(duties []d.Duty, toFind string) (string, error) {
	return findInDuties(duties, "summary", toFind)
}

func findInPriorities(duties []d.Duty, toFind string) (string, error) {
	return findInDuties(duties, "priority", toFind)
}

func findInStatus(duties []d.Duty, toFind string) (string, error) {
	return findInDuties(duties, "status", toFind)
}

func findInBody(duties []d.Duty, toFind string) (string, error) {
	return findInDuties(duties, "body", toFind)
}

func findInAny(duties []d.Duty, toFind string) (string, error) {
	return findInDuties(duties, "any", toFind)
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
