package main

import (
	"blog/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedFollows, err := apiCfg.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't create feed %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfig.db.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't get feed %s", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollowIDstring := chi.URLParam(r, "feedfollowsID")

	feedfollowID, err := uuid.Parse(feedfollowIDstring)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid parsing feed follow ID")
		return
	}

	err = apiConfig.db.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't delete feed %s", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})

}

// func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
// 	feeds, err := apiCfg.db.GetFeeds(r.Context())
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("couldn't get feeds %s", err))
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, databaseFeedsToFeeds(feeds))
// }
