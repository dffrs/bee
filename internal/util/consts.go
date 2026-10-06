package internal

const (
	BeeDir = ".bee"
	BeeExt = ".md"
)

type Color = string

const (
	Reset   Color = "\033[0m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Magenta Color = "\033[35m"
	Gray    Color = "\033[37m"
)
