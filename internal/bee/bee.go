package internal

import (
	"flag"
	"fmt"

	cmd "bee/internal/commands"
	ut "bee/internal/util"
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
		"purge":   cmd.Command.Purge,
		"help":    cmd.Command.Help,
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
		helpCmd := b.commands["help"]

		help, err := helpCmd(nil, args)
		if err != nil {
			return err
		}

		return fmt.Errorf(ut.Yellow+"[Bee]\n"+ut.Reset+"command not recognized.\n\n%s", help)
	}

	feedBack, err := toRun(nil, args)
	if err != nil {
		return fmt.Errorf(ut.Yellow+"[Bee]\n"+ut.Reset+"'%s' error: %w", flg, err)
	}

	fmt.Printf(ut.Yellow+"[Bee]\n"+ut.Reset+"%s", feedBack)

	return nil
}
