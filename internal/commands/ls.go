package internal

import (
	"os"
	"path/filepath"
	"strings"

	d "bee/internal/duty"
	ut "bee/internal/util"
)

func findMarkdownFiles(path string, duties []d.Duty) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// TODO: have const for this
		if filepath.Ext(info.Name()) != ".md" {
			return nil
		}

		duty := new(d.Duty)
		if err = d.BuildFromMarkdown(path, duty); err != nil {
			return err
		}

		duties = append(duties, *duty)

		return nil
	})

	return err
}

func (cm Command) LS(args []string) (string, error) {
	path, err := ut.TryGetBeeDir()
	if err != nil {
		return "", err
	}

	duties := make([]d.Duty, 0)
	if err := findMarkdownFiles(path, duties); err != nil {
		return "", err
	}

	// no duties found
	if len(duties) == 0 {
		return "not duties found. Run 'bee new <summary>' to create a new duty\n", nil
	}

	builder := strings.Builder{}

	builder.WriteString("\n")
	for _, d := range duties {
		builder.WriteString(d.ToString())
		builder.WriteString("\n")
	}

	return builder.String(), nil
}
