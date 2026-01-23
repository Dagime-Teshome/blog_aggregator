package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Dagime-Teshome/blog_aggregator/internal/database"
)

func Browse(s *State, cmd Command, user database.User) error {
	queryLimit := int32(2)
	if len(cmd.Args) > 0 {
		if n, err := strconv.Atoi(cmd.Args[0]); err != nil || n <= 0 {
			return fmt.Errorf("limit must be a positive number")
		} else {
			queryLimit = int32(n)
		}
	}
	getParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  queryLimit,
	}
	postList, err := s.Db.GetPostsForUser(context.Background(), getParams)
	if err != nil {
		return fmt.Errorf("error getting posts for user %s err:%v", user.Name, err)
	}
	if len(postList) == 0 {
		fmt.Println("no posts in list")
	}
	for i, posts := range postList {
		fmt.Printf("%v : %s \n", i, posts.Title)
	}
	return nil
}
