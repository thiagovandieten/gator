package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"
	"time"

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
	validURL := isValidUrl(cmd.Args[1])
	if !validURL {
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

func handlerFeeds(s *state, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feeds) < 1 {
		fmt.Printf("No feeds saved")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 1, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Name\tURL\tUsername")
	for _, feed := range feeds {
		fmt.Fprintf(w, "%s\t%s\t%s\n", feed.Name, feed.Url, feed.Username)
	}
	w.Flush()
	return nil
}

func handlerFollow(s *state, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("no url given to follow")
	}

	validURL := isValidUrl(cmd.Args[0])
	if !validURL {
		return errors.New("not a valid url provided")
	}

	feed, err := s.db.SearchFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
	}

	ffParams := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedfollow, err := s.db.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s %s", feed.Name, user.Name)

	return nil
}

func isValidUrl(paramURL string) bool {
	url, err := url.Parse(paramURL)
	if err != nil || url.Scheme == "" || url.Host == "" {
		return false
	}
	return true
}
