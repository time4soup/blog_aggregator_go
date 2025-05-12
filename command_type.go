package main

import (
	"fmt"
)

func NewCommands() commands {
	return commands{map[string]func(*state, command) error{}}
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (cmds *commands) run(s *state, cmd command) error {
	if f, ok := cmds.registry[cmd.name]; ok {
		err := f(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("commmand not found: %s", cmd.name)
	}

	return nil
}

func (cmds *commands) register(name string, f func(*state, command) error) {
	cmds.registry[name] = f
}

type command struct {
	name string
	args []string
}
