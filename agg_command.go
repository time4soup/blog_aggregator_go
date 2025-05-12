package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/time4soup/blog_aggregator_go/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command 'agg' takes one argument")
	}
	time_between_reqs := cmd.args[0]

	reqFreq, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)

	ticker := time.NewTicker(reqFreq)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	params := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Fetching feed: %s\n", feed.Name)
	rss_feed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range rss_feed.Channel.Item {
		savePost(s, item, feed.ID)
	}
}

func savePost(s *state, item RSSItem, feedId uuid.UUID) {
	_, err := s.db.GetPostByUrl(context.Background(), item.Link)
	if err == nil {
		return
	}

	pubTime, err := parseVariousTimes(item.PubDate)
	if err != nil {
		return
	}
	params := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       item.Title,
		Url:         item.Link,
		Description: item.Description,
		PublishedAt: pubTime,
		FeedID:      feedId,
	}
	_, err = s.db.CreatePost(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
}

func parseVariousTimes(timeString string) (time.Time, error) {
	t, err := time.Parse(time.RFC1123Z, timeString)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}
