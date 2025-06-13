package main

import (
	"context"
	"fmt"

	"github.com/thiagovandieten/gator/internal/config"
)

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
