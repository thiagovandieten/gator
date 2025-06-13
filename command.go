package main

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandHandlers map[string]func(*state, Command) error
}

func (c *Commands) run(s *state, cmd Command) error {
	if c.CommandHandlers[cmd.Name] == nil {
		return errors.New("command not found")
	}

	err := c.CommandHandlers[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, f func(*state, Command) error) error {
	if c.CommandHandlers == nil {
		return errors.New("commandHandlers map uninitialized")
	}
	c.CommandHandlers[name] = f
	return nil
}

// registerAll registers multiple command handlers at once
func (c *Commands) registerAll(handlers map[string]func(*state, Command) error) error {
	if c.CommandHandlers == nil {
		return errors.New("commandHandlers map uninitialized")
	}

	for name, handler := range handlers {
		c.CommandHandlers[name] = handler
	}

	return nil
}
