package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func( *State, Command) error {
	return func(s *State, cmd Command) error {
		 user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		 if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		 }
		 return handler(s, cmd, user)
	}
}