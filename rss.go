package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}

	defer resp.Body.Close()

	dat, err :=io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}

	rssfeed := RssFeed{}
	err = xml.Unmarshal(dat, &rssfeed)
	if err != nil {
		return RssFeed{}, err
	}

	return rssfeed, nil

}
