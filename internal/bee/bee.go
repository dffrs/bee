package internal

import (
	"errors"
	"flag"
	"fmt"

	cmd "bee/internal/commands"
)

type Cmd func(st cmd.Command, args []string) (string, error)

type Bee struct {
	commands map[string]Cmd
}

func NewBee() *Bee {
	flag.Parse()

	commands := map[string]Cmd{
		"start": cmd.Command.Start,
	}

	return &Bee{commands: commands}
}

func (b *Bee) Run() error {
	args := flag.Args()

	flg := args[0]
	vals := args[1:]

	toRun, ok := b.commands[flg]
	if !ok {
		return errors.New("[Bee]: command not recognized")
	}

	feedBack, err := toRun(nil, vals)
	if err != nil {
		return fmt.Errorf("[Bee]: '%s' error: %w", flg, err)
	}

	fmt.Printf("[Bee]: %s", feedBack)

	return nil
}
