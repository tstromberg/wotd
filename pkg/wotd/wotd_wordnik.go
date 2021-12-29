package wotd

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Wordnik returns words-of-the-day from wordnik.com
func Wordnik(ctx context.Context) (*Result, error) {
	url := "https://wordnik.com/word-of-the-day"
	doc, err := fetch(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}

	r := &Result{Kind: "wordnik.com", URL: url}
	doc.Find("#wotd h1").Each(func(i int, s *goquery.Selection) {
		r.Word = s.Text()
	})

	doc.Find("div.guts").Each(func(i int, s *goquery.Selection) {

		s.Find("h3.source").Each(func(i int, s *goquery.Selection) {
			r.Definitions = append(r.Definitions, Definition{Source: s.Text()})
		})

		count := 0
		s.Find("ul").Each(func(i int, ul *goquery.Selection) {
			ul.Find("li").Each(func(i int, li *goquery.Selection) {
				_, suffix, _ := strings.Cut(li.Text(), " ")

				ps := PartOfSpeech{
					Text: suffix,
				}
				li.Find("abbr").Each(func(i int, abbr *goquery.Selection) {
					ps.Kind = abbr.Text()
				})
				r.Definitions[count].Parts = append(r.Definitions[count].Parts, ps)
			})
			count++
		})
	})

	doc.Find("li.exampleItem").Each(func(i int, li *goquery.Selection) {
		ex := Example{}
		li.Find("p.text").Each(func(i int, p *goquery.Selection) {
			ex.Text = p.Text()
		})
		li.Find("p.source a").Each(func(i int, a *goquery.Selection) {
			ex.Source = a.Text()
			ex.URL = a.AttrOr("href", "")
		})
		r.Examples = append(r.Examples, ex)
	})

	return r, nil
}
