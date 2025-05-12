package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	if len(cmd.args) != 0 && len(cmd.args) != 1 {
		return fmt.Errorf("command 'browse' takes zero or one arguments")
	}
	var limit int
	if len(cmd.args) == 0 {
		limit = 2
	} else {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
	}

	params := database.GetPostsForUserParams{
		Name:  s.config.CurrentUserName,
		Limit: int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s\n", post.Description)
		fmt.Printf("%s\n\n", post.PublishedAt)
	}

	return nil
}
