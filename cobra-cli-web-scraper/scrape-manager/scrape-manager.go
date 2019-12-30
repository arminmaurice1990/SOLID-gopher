package scrape_manager

import (
	"cobra-cli-web-scraper/saver"
	"cobra-cli-web-scraper/scraper"
	"fmt"
)

type ScrapeManager interface {
	Run(urls []string)
}

type scrapemanager struct {
	scraper scraper.Scraper
	saver saver.Saver
}

func NewScrapeManager(scraper scraper.Scraper, saver saver.Saver) *scrapemanager {
	return &scrapemanager{scraper:scraper, saver:saver}
}

func (s *scrapemanager) Run(urls []string) {
	for _,url:= range urls {
		sitereader, err := s.scraper.Scrape(url)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = s.saver.Save(url, sitereader)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}