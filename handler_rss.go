package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/thiagovandieten/gator/internal/database"
)

func handlerAgg(s *state, cmd Command) error {
	feed, err := fetchFeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%-v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("not enough arguments provided to add feed")
	}

	// Check if the URL is valid
	url, err := url.Parse(cmd.Args[1])
	if err != nil || url.Scheme == "" || url.Host == "" {
		return errors.New("the second argument is not a valid URL")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
	}

	// feed, err := fetchFeed(cmd.Args[1])
	// if err != nil {
	// 	return err
	// }

	params := database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: user.ID,
	}

	result, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("%-v\n", result)

	return nil
}
