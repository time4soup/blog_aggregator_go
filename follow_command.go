package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'follow' takes one argument")
	}

	Url := cmd.args[0]
	FeedId, err := s.db.GetFeedIdFromUrl(context.Background(), Url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    FeedId,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' followed '%s'.\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}
