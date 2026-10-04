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
		"start":   cmd.Command.Start,
		"version": cmd.Command.Version,
		"new":     cmd.Command.New,
		"ls":      cmd.Command.LS,
	}

	return &Bee{commands: commands}
}

func (b *Bee) Run() error {
	args := flag.Args()

	var flg string
	if len(args) == 0 {
		flg = ""
	} else {
		flg = args[0]
	}

	toRun, ok := b.commands[flg]
	if !ok {
		return errors.New("[Bee]\ncommand not recognized")
	}

	feedBack, err := toRun(nil, args)
	if err != nil {
		return fmt.Errorf("[Bee]\n'%s' error: %w", flg, err)
	}

	fmt.Printf("[Bee]\n%s", feedBack)

	return nil
}
