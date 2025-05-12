package main

import (
	"context"
	"fmt"

	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'unfollow' takes one argument")
	}

	params := database.DeleteFeedFollowParams{
		Url:    cmd.args[0],
		UserID: user.ID,
	}
	err := s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' unfollowed '%s'\n", user.Name, cmd.args[0])
	return nil
}
