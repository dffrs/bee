package internal

func (cm Command) LS(args []string) (string, error) {
	path, err := tryGetBeeDir()
	if err != nil {
		return "", err
	}

	_ = path

	return "", nil
}
