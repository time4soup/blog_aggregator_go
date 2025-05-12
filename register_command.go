package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'register' takes one command")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("duplicate user: %s", cmd.args[0])
	}

	_, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User created: %s\n", cmd.args[0])
	return nil
}
