package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"log"
	"os"
)

func handlerAddFeed (s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
	  log.Fatal("too few arguments, usage: addfeed <name> <url>")
		os.Exit(1)
	}

  name := cmd.Args[0]
	url := cmd.Args[1]
	ctx := context.Background()
  
	newFeed := database.CreateFeedParams{Name: name, Url: url, UserID: user.ID}
	rssFeed, err := s.db.CreateFeed(ctx, newFeed )
	if err != nil {
		return fmt.Errorf("error creating new feed: %v", err)
	}
	followfeed := database.CreateFeedFollowParams{UserID: user.ID, FeedID: rssFeed.ID}
	_ , err = s.db.CreateFeedFollow(ctx, followfeed)
	if err != nil {
		return fmt.Errorf("error following feed after creating feed: %v", err)
	}

	fmt.Println("Feed was created")
	fmt.Printf("Feed DATA: \n")
	fmt.Printf("{\n  id: %v,\n  name: %v,\n  url: %v,\n  user_id: %v,\n  created_at: %v,\n  updated_at: %v\n}\n", rssFeed.ID, rssFeed.Name, rssFeed.Url, rssFeed.UserID, rssFeed.CreatedAt, rssFeed.UpdatedAt)
	return nil
}