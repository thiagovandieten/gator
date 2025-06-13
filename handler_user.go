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

func handlerGetAllUsers(s *state, cmd Command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {

		userStr := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.Username {
			userStr += " (current)"
		}
		fmt.Println(userStr)
	}

	return nil
}
