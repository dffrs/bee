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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO: build duty struct here
		fmt.Println(scanner.Text())
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
