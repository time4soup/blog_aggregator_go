package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	res_data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rss_feed := RSSFeed{}
	xml.Unmarshal(res_data, &rss_feed)

	rss_feed.Channel.Title = html.UnescapeString(rss_feed.Channel.Title)
	rss_feed.Channel.Description = html.UnescapeString(rss_feed.Channel.Description)
	for i, item := range rss_feed.Channel.Item {
		rss_feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rss_feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &rss_feed, nil
}
