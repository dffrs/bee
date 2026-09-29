package internal

import d "bee/internal/duty"

func (cm Command) LS(args []string) (string, error) {
	path, err := tryGetBeeDir()
	if err != nil {
		return "", err
	}

	// TODO: path is the bee dir, not a .md file.
	var duty *d.Duty
	if err = d.BuildFromMarkdown(path, duty); err != nil {
		return "", err
	}

	return duty.ToString(), nil
}
