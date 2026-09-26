package internal

import (
	"errors"
	"flag"
	"fmt"

	cmd "bee/internal/commands"
)

type Bee struct {
	commands []cmd.Command
}

func NewBee() *Bee {
	flag.Parse()
	commands := []cmd.Command{}

	return &Bee{commands: commands}
}

func (b *Bee) Run() error {
	args := flag.Args()

	flg := args[0]
	vals := args[1:]

	var toRun cmd.Command
	for _, cmd := range b.commands {
		if cmd.Name == flg {
			toRun = cmd
			break
		}
	}

	if toRun.Name == "" {
		return errors.New("command not recognized")
	}

	feedBack, err := toRun.Run(vals)
	if err != nil {
		return fmt.Errorf("[%s] error: %w", toRun.Name, err)
	}

	fmt.Printf("%s", feedBack)

	return nil
}
