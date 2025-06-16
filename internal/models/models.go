package models

import (
	"encoding/xml"
	"time"
)

// RssFeed handles the content according to RSS 2.0 specification
type RssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	Title string    `xml:"title"`
	Items []RssItem `xml:"item"`
}

type RssItem struct {
	MediaTitle string `xml:"http://search.yahoo.com/mrss/ title"`
	Title      string `xml:"title"`
	Link       string `xml:"link"`
	PubDate    string `xml:"pubDate"`
}

type Item struct {
	Title       string
	Link        string
	PublishedAt time.Time
}

type Feed struct {
	Title string
	Items []*Item
	Order int
}

type Page struct {
	Title       string
	LastUpdated time.Time
	Feeds       []*Feed
}
