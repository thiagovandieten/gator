package main

import (
	"context"

	"github.com/thiagovandieten/gator/internal/database"
)

func middlewareLoggedin(handler func(s *state, cmd Command, user database.User) error) func(*state, Command) error {
	return func(s *state, c Command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
