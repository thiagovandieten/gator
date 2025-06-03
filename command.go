package main

import (
	"errors"
	"fmt"

	"github.com/thiagovandieten/gator/internal/config"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	Names map[string]func(*state, command)
}

func (c *commands) run(s *state, cmd command) error {
	if c.Names[cmd.Name] != nil {
		c.Names[cmd.Name](s, cmd)
		return nil
	}
	return errors.New("command not found")
}

func (c *commands) register(name string, f func(*state, command)) error {

	return nil
}

func handlerLogin(s *state, cmd command) error {
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
