package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required; usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]
	ctx := context.Background()

	if _, err := s.db.GetUser(ctx, username); err != nil {
		log.Fatalf("user %q does not exist", username)
		os.Exit(1)
	}
	err := s.cfg.SetUser(username)
  if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("The user has been set.")
	return nil
}