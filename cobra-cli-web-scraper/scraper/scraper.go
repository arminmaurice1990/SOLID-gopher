package scraper

import (
	"io"
	"net/http"
)

type Scraper interface {
	Scrape(url string) (io.Reader, error)
}

type scraper struct{}

func NewScraper() *scraper {
	return &scraper{}
}

func (s *scraper) Scrape(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}
