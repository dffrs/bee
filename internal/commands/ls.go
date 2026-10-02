package internal

import (
	"os"
	"path/filepath"
	"strings"

	d "bee/internal/duty"
)

func (cm Command) LS(args []string) (string, error) {
	path, err := tryGetBeeDir()
	if err != nil {
		return "", err
	}

	duties := make([]d.Duty, 0)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(info.Name()) != ".md" {
			return nil
		}

		var duty *d.Duty
		if err = d.BuildFromMarkdown(path, duty); err != nil {
			return err
		}

		duties = append(duties, *duty)

		return nil
	})
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	for _, d := range duties {
		builder.WriteString(d.ToString())
	}

	return builder.String(), nil
}
