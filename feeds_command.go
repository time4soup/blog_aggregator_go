package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command 'feeds' takes no arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds added")
	}
	for _, feed := range feeds {
		name, err := s.db.GetUserFromId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("name: %s\n", feed.Name)
		fmt.Printf("url: %s\n", feed.Url)
		fmt.Printf("user: %s\n\n", name)
	}

	return nil
}
