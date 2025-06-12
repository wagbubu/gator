package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *State, cmd Command) error {
	ctx := context.Background()
  followingfeeds, err := s.db.GetFeedFollowsForUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting all following feeds: %v", err)
	}
	fmt.Println("--- LIST OF ALL FOLLOWING FEEDS ---")
	for idx, feed := range followingfeeds {
		fmt.Println("")
		fmt.Printf("-[ FEED %d ]------------\n", idx + 1)
		fmt.Printf("name: %q\n", feed.FeedName)
		fmt.Println("")
	}
	return nil
}