package main

import (
	"blog-aggregator/internal/config"
	"blog-aggregator/internal/database"
	"errors"
)

type State struct {
	db *database.Queries
	cfg *config.Config
}

type Command struct {
	Name string;
	Args []string;
}

type Commands struct {
	handlers map[string]func(*State, Command) error
} 

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.handlers[cmd.Name]
	if !ok {
		return errors.New("error: command not found")
	}
	return f(s, cmd)
}

func (c *Commands) register(Name string, f func(*State, Command) error) {
	c.handlers[Name] = f  
}