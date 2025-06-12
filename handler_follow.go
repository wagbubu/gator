package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("no arguments found; usage: follow <feed-url>")
	}
	ctx := context.Background()
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("error getting feeds by url: %v", err)
	}
	
	feedfollow := database.CreateFeedFollowParams{UserID: user.ID, FeedID: feed.ID}
	newfeedfollow, err := s.db.CreateFeedFollow(ctx, feedfollow)	
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}
	fmt.Println("Creating feed follow success!")
	fmt.Printf("username: %q\n", newfeedfollow.UserName)
	fmt.Printf("feed name: %q\n", newfeedfollow.FeedName)
	return nil
}