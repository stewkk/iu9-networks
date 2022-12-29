package kommersant

import (
	"github.com/SlyMarbo/rss"
)

func ParseFeed() ([]NewsEntry, error) {
	feed, err := rss.Fetch("https://www.kommersant.ru/RSS/main.xml")
	if err != nil {
		return nil, err
	}

	var news []NewsEntry
	for _, entry := range feed.Items {
		news = append(news, NewsEntry{
			Id:          entry.ID,
			Link:        entry.Link,
			Category:    entry.Categories[0],
			Title:       entry.Title,
			PubDate:     entry.Date.String(),
			Description: entry.Summary,
		})
	}
	return news, nil
}
