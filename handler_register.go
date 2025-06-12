package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username required; usage: %s <username>", cmd.Name)
	}
	ctx := context.Background()
	username := cmd.Args[0]
	
	if _ , err := s.db.GetUser(ctx, username); err == nil {
		 log.Fatalf("username %q already exists\n", username)
		 os.Exit(1)
	}
	user := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: username}
	s.db.CreateUser(ctx, user)

	s.cfg.SetUser(username)
	fmt.Println("Username was created")
	fmt.Printf("USER DATA: \n")
	fmt.Printf("{\n  id: %v,\n  name: %v,\n  created_at: %v,\n  updated_at: %v\n}\n", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}