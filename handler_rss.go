package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
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

func handlerAddFeed(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("not enough arguments provided to add feed")
	}

	// Check if the URL is valid
	validURL := isValidUrl(cmd.Args[1])
	if !validURL {
		return errors.New("the second argument is not a valid URL")
	}

	params := database.CreateFeedParams{
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	result, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("%-v\n", result)

	params_feedfollow := database.CreateFeedFollowParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    result.ID,
	}
	s.db.CreateFeedFollow(context.Background(), params_feedfollow)

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

	// w := tabwriter.NewWriter(os.Stdout, 2, 1, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintln(s.tab, "Name\tURL\tUsername")
	for _, feed := range feeds {
		fmt.Fprintf(s.tab, "%s\t%s\t%s\n", feed.Name, feed.Url, feed.Username)
	}
	s.tab.Flush()
	return nil
}

func handlerFollow(s *state, cmd Command, user database.User) error {
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
	fmt.Fprintln(s.tab, "Username\tFeedname\tID\tCreatedAt\tUpdatedAt\tUserID\tFeedID")
	fmt.Fprintf(
		s.tab,
		"%s\t%s\t%d\t%s\t%s\t%s\t%d\n",
		feedfollow.Username,
		feedfollow.Feedname,
		feedfollow.ID,
		feedfollow.CreatedAt,
		feedfollow.UpdatedAt,
		feedfollow.UserID,
		feedfollow.FeedID,
	)

	s.tab.Flush()
	return nil
}

func handlerFollowing(s *state, cmd Command) error {

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.Username)
	if err != nil {
		return err
	}
	fmt.Printf("%s is following these feeds:\n", s.cfg.Username)
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed)
	}
	return nil
}

func handlerUnfollowing(s *state, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("no feed url provided")
	}

	if !isValidUrl(cmd.Args[0]) {
		return errors.New("not a valid url provided")
	}

	params := database.DeleteFeedFollowWithFeedURLParams{
		Url:    cmd.Args[0],
		UserID: user.ID,
	}
	s.db.DeleteFeedFollowWithFeedURL(context.Background(), params)
	return nil
}

func scrapeFeeds(s *state, cmd Command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(ctx.Background(), user.Name)
	if err != nil {
		return err
	}
	if len(feeds) < 1 {
		return errors.New("no feeds found for this user")
	}
	return nil
}

func isValidUrl(paramURL string) bool {
	url, err := url.Parse(paramURL)
	if err != nil || url.Scheme == "" || url.Host == "" {
		return false
	}
	return true
}
