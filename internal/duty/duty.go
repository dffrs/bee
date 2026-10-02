package internal

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func NewDuty(summary string, priority int, body string) *Duty {
	return &Duty{Summary: summary, Priority: priority, Body: body}
}

func BuildFromMarkdown(path string, duty *Duty) error {
	if filepath.Ext(path) != ".md" {
		return errors.New("file extension not supported")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	linesRead := 0
	hasReadSummary := false
	hasReadPriority := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// must start with '## Summary'
		if linesRead == 0 && !strings.EqualFold(strings.TrimSpace(line), summaryHeader) {
			return fmt.Errorf("first line must have a '%s' header", summaryHeader)
		}

		// so that we don't append this to body
		if strings.EqualFold(line, summaryHeader) {
			continue
		}

		// are we done with summary section ?
		if strings.EqualFold(line, priorityHeader) {
			hasReadSummary = true
			continue
		}

		// are we done with priority section ?
		if strings.EqualFold(line, bodyHeader) {
			hasReadPriority = true
			continue
		}

		if !hasReadSummary {
			duty.Summary += line
			continue
		}

		if !hasReadPriority {
			prio, err := strconv.Atoi(line)
			if err != nil {
				return fmt.Errorf("failed to get priority: got '%s'. error: %w", line, err)
			}

			duty.Priority = prio
			continue
		}

		duty.Body += line

		linesRead++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

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
	content.WriteString("")

	return content.String()
}
