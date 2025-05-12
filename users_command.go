package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command 'users' takes no arguments")
	}

	usersList, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(usersList) == 0 {
		fmt.Println("No users registered")
		return nil
	}

	currentUser := s.config.CurrentUserName
	for _, u := range usersList {
		fmt.Printf("* %s", u)
		if u == currentUser {
			fmt.Print(" (current)")
		}
		fmt.Print("\n")
	}

	return nil
}
