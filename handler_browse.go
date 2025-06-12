package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if userLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = userLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	fmt.Println("--- LIST OF ALL POSTS ---")
	for idx, post := range posts {
		fmt.Println("")
		fmt.Printf("-[ POST %d ]------------\n", idx+1)
		fmt.Printf("Published: %q\n", post.PublishedAt.Time)
		fmt.Printf("Title: %q\n", post.Title)
		fmt.Printf("Description: %q\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
	}
	return nil
}