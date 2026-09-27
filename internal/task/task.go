package internal

import (
	"strconv"
	"strings"
)

type Task struct {
	Summary  string
	Priority int
	Body     string
}

func NewTask(summary string, priority int, body string) *Task {
	return &Task{Summary: summary, Priority: priority, Body: body}
}

func (t *Task) ToMarkdown() string {
	content := strings.Builder{}

	content.WriteString("## Summary")
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(t.Summary)
	content.WriteString("\n")
	content.WriteString("\n")

	content.WriteString("## Priority")
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(strconv.Itoa(t.Priority))
	content.WriteString("\n")
	content.WriteString("\n")

	content.WriteString("## Body")
	content.WriteString("\n")
	content.WriteString("\n")
	content.WriteString(t.Body)
	content.WriteString("\n")

	return content.String()
}
