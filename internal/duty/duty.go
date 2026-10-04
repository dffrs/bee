package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	ut "bee/internal/util"
)

const (
	summaryHeader  string = "## Summary"
	priorityHeader string = "## Priority"
	bodyHeader     string = "## Body"
)

type Duty struct {
	Summary  string
	Priority int
	Body     string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func splitSections(content string) (summary, priority, body string, err error) {
	headers := []string{summaryHeader, priorityHeader, bodyHeader}

	var (
		sections = make([]string, len(headers))
		opened   = -1
	)

	for _, raw := range strings.Split(content, "\n") {
		line := strings.TrimSpace(raw)

		idx := slices.IndexFunc(headers, func(header string) bool {
			return strings.EqualFold(line, header)
		})

		if idx >= 0 {
			if idx != opened+1 {
				return "", "", "", fmt.Errorf("expected the '%s' header, got '%s'", headers[opened+1], line)
			}

			opened = idx
			continue
		}

		if opened < 0 {
			return "", "", "", fmt.Errorf("first line must have a '%s' header", summaryHeader)
		}

		sections[opened] += line + "\n"
	}

	if opened != len(headers)-1 {
		return "", "", "", fmt.Errorf("'%s' header not found", headers[opened+1])
	}

	return strings.TrimSpace(sections[0]), strings.TrimSpace(sections[1]), strings.TrimSpace(sections[2]), nil
}

func parseMarkdown(content string) (Duty, error) {
	summary, priorityText, body, err := splitSections(content)
	if err != nil {
		return Duty{}, err
	}

	priority, err := strconv.Atoi(priorityText)
	if err != nil {
		return Duty{}, fmt.Errorf("failed to get priority: got '%s'. error: %w", priorityText, err)
	}

	return Duty{Summary: summary, Priority: priority, Body: body}, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewDuty(summary string, priority int, body string) *Duty {
	return &Duty{Summary: summary, Priority: priority, Body: body}
}

func BuildFromMarkdown(path string, duty *Duty) error {
	if filepath.Ext(path) != ut.BeeExt {
		return errors.New("file extension not supported")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read duty file (%s): %w", path, err)
	}

	parsed, err := parseMarkdown(string(content))
	if err != nil {
		return fmt.Errorf("invalid duty file (%s): %w", path, err)
	}

	*duty = parsed

	return nil
}

func (t *Duty) ToMarkdown() string {
	content := strings.Builder{}

	content.WriteString(summaryHeader)
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(t.Summary)
	content.WriteString("\n")
	content.WriteString("\n")

	content.WriteString(priorityHeader)
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(strconv.Itoa(t.Priority))
	content.WriteString("\n")
	content.WriteString("\n")

	content.WriteString(bodyHeader)
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(t.Body)
	content.WriteString("\n")

	return content.String()
}

func (t *Duty) ToString() string {
	content := strings.Builder{}

	content.WriteString("Priority: ")
	content.WriteString(strconv.Itoa(t.Priority))
	content.WriteString(" | ")

	content.WriteString("Summary: ")
	content.WriteString(t.Summary)

	return content.String()
}
