package internal

import "fmt"

// Version is the released version, injected at build time with
// -ldflags "-X bee/internal/commands.Version=<tag>".
var Version = "dev"

func (cm Command) Version(args []string) (string, error) {
	return fmt.Sprintf("%s\n", Version), nil
}
