package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command 'reset' take no arguments")
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	err = s.db.ResetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Users and Feeds tables reset.")
	return nil
}
