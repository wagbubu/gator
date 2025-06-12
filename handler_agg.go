package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("no arguments found; usage: agg <request interval>")
	}
	interval := cmd.Args[0]
	timeBetweenReqs, err := time.ParseDuration(interval)
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}

	fmt.Printf("collecting feeds every %s\n", interval)
	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}	
}

func scrapeFeeds(s *State)  {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("error getting next feed to fetch: %v", err)
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v", err)
	}

	freshFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("error fetching next feed: %v", err)
	}

	for _, item := range freshFeed.Channel.Items {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		newPost := database.CreatePostParams{
			Title: item.Title, 
			Url: item.Link, 
			FeedID: feed.ID, 
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
		}
		_, err := s.db.CreatePost(ctx, newPost)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}  
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(freshFeed.Channel.Items))
}
