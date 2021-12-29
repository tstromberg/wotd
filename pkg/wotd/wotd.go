package wotd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Kind string
	URL  string

	Word        string
	Definitions []Definition
	Examples    []Example
}

type Definition struct {
	Source string
	Parts  []PartOfSpeech
}

type PartOfSpeech struct {
	Kind string
	Text string
}

type Example struct {
	Source string
	URL    string
	Text   string
}

func fetch(ctx context.Context, url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s status code %d", url, res.StatusCode)
	}

	return goquery.NewDocumentFromReader(res.Body)
}

// All returns words-of-the-day from all known sources
func All(ctx context.Context) ([]*Result, error) {
	rs := []*Result{}
	r, err := Wordnik(ctx)
	if err != nil {
		return rs, fmt.Errorf("wordnik: %w", err)
	}
	rs = append(rs, r)

	r, err = Dictionary(ctx)
	if err != nil {
		return rs, fmt.Errorf("dictionary: %w", err)
	}
	rs = append(rs, r)

	return rs, nil
}
