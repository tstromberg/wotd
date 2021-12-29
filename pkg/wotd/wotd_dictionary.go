package wotd

import (
	"context"
	"fmt"
)

// Dictionary returns words-of-the-day from dictionary.com
func Dictionary(ctx context.Context) (*Result, error) {
	url := "https://www.dictionary.com/e/word-of-the-day/"
	doc, err := fetch(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	block := doc.Find("div.otd-item-headword__pos").First()

	r := &Result{
		Kind: "dictionary.com",
		URL:  url,
		Word: doc.Find("h1.js-fit-text").First().Text(),
		Definitions: []Definition{
			{
				Source: "unknown",
				Parts: []PartOfSpeech{
					{
						Kind: block.Find("span.luna-pos").First().Text(),
						Text: block.Find("p").Last().Text(),
					},
				},
			},
		},
	}

	return r, nil
}
