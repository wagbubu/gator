package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *State, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds from database: %v", err)
	}
	fmt.Println("--- LIST OF ALL FEEDS ---")
	for idx, feed := range feeds {
		fmt.Println("")
		fmt.Printf("-[ FEED %d ]------------\n", idx+1)
		fmt.Printf("name: %q\n", feed.Name)
		fmt.Printf("url: %q\n", feed.Url)
		user, err := s.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user by id referenced in feeds: %v", err)
		}
		fmt.Printf("created_by: %q\n", user.Name)
		fmt.Println("")
	}
	return nil
}
