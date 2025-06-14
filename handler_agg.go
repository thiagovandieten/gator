package main

import "fmt"

func handlerAgg(s *state, cmd Command) error {
	feed, err := fetchFeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%-v\n", feed)
	return nil
}
