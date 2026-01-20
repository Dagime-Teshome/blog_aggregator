package feed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", "gator")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rssFeedByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rssFeed RSSFeed
	err = xml.Unmarshal(rssFeedByte, &rssFeed)
	if err != nil {
		return nil, fmt.Errorf("error marshalling:%v", err)
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, value := range rssFeed.Channel.Item {
		value.Title = html.UnescapeString(value.Title)
		value.Description = html.UnescapeString(value.Description)
	}
	return &rssFeed, nil
}
