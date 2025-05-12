package main

import (
	"context"
	"fmt"

	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command 'folliwng' takes no arguments")
	}

	followingFeeds, err := s.db.GetFeedFollowForUser(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	if len(followingFeeds) == 0 {
		fmt.Printf("User '%s' has no followed feeds.\n", currentUser.Name)
		return nil
	}
	fmt.Printf("User '%s' following:\n", currentUser.Name)
	for _, followingFeed := range followingFeeds {
		fmt.Printf("* %s\n", followingFeed.FeedName)
	}

	return nil
}
