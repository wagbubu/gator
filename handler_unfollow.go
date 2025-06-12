package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
)

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("error no argument found; usage: unfollow <feed-url>")
	}
	feedURL := cmd.Args[0]
	ctx := context.Background()
	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v", err)
	}
  feedToUnfollow := database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID }
	err = s.db.DeleteFeedFollow(ctx, feedToUnfollow)
  if err != nil {
		return fmt.Errorf("error unfollowing feed")
	}
	fmt.Printf("You have unfollowed the feed %q successfully!", feed.Name)
	return nil
} 