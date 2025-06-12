package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thiagovandieten/gator/internal/config"
	"github.com/thiagovandieten/gator/internal/database"
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

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.Args) == 0 || cmd.Args[0] == "" {
		return errors.New("no username was given")
	}

	username := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = config.SetUser(username, s.cfg)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s has been set!\n", username)
	return nil
}

func handlerRegister(s *state, cmd Command) error {
	if len(cmd.Args) <= 0 || cmd.Args[0] == "" {
		return errors.New("no username was given")
	}
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	result, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been created!\n", result.Name)
	fmt.Printf("DEBUG: User info:\n%-v\n", result)
	config.SetUser(result.Name, s.cfg)

	return nil
}

func handlerReset(s *state, cmd Command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("Something went wrong attempting to remove all users")
		return err
	}
	fmt.Println("All users deleted")
	config.SetUser("", s.cfg)
	return nil
}

func handlerGetAllUsers(s *state, cmd Command) error {
	return nil
}
