package main

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	ctx := context.Background()
	if err := s.db.Reset(ctx); err != nil {
		return fmt.Errorf("error resetting database: %v", err)
	}
	fmt.Println("Database reset was successful.")
	return nil
}