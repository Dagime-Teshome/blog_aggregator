package cli

import (
	"context"
	"fmt"
	feed "go/gator/internal/rss_feed"
)

func Agg(s *State, cmd Command) error {
	ctx := context.Background()
	rssUrl := "https://www.wagslane.dev/index.xml"
	rssfeed, err := feed.FetchFeed(ctx, rssUrl)
	if err != nil {
		return err
	}
	fmt.Println(rssfeed)
	return nil
}
