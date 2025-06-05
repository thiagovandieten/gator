package main

import (
	"errors"
	"fmt"

	"github.com/thiagovandieten/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Names map[string]func(*state, Command)
}

func (c *Commands) run(s *state, cmd Command) error {
	if c.Names[cmd.Name] != nil {
		c.Names[cmd.Name](s, cmd)
		return nil
	}
	return errors.New("command not found")
}

func (c *Commands) register(name string, f func(*state, Command)) error {
	if c.Names == nil {
		c.Names = make(map[string]func(*state, Command))
	}
	c.Names[name] = f
	return nil
}

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no username was given")
	}

	username := cmd.Args[0]
	err := config.SetUser(username, *s.cfg)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s has been set!\n", username)
	return nil
}
