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
	rssfeedByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rssfeed RSSFeed
	err = xml.Unmarshal(rssfeedByte, &rssfeed)
	if err != nil {
		return nil, fmt.Errorf("error marshalling:%v", err)
	}
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
	for _, value := range rssfeed.Channel.Item {
		value.Title = html.UnescapeString(value.Title)
		value.Description = html.UnescapeString(value.Description)
	}
	return &rssfeed, nil
}
