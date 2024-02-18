package main

import (
	"blog/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbuser database.User) User {
	return User{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}

}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

func databaseFeedsToFeeds(dbfeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbfeeds))
	for i, dbfeed := range dbfeeds {
		feeds[i] = databaseFeedtoFeed(dbfeed)
	}
	return feeds
}

type feedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedFollowstoFeedFollows(dbfeedFollows database.FeedFollow) feedFollow {
	return feedFollow{
		ID:        dbfeedFollows.ID,
		CreatedAt: dbfeedFollows.CreatedAt,
		UpdatedAt: dbfeedFollows.UpdatedAt,
		FeedID:    dbfeedFollows.FeedID,
		UserID:    dbfeedFollows.UserID,
	}
}


func databaseFeedFollowToFeedFollow(dbfeedFollows []database.FeedFollow) []feedFollow {
	feedFollows := make([]feedFollow, len(dbfeedFollows))
	for i, dbfeedFollow := range dbfeedFollows {
		feedFollows[i] = databaseFeedFollowstoFeedFollows(dbfeedFollow)
	}
	return feedFollows
}