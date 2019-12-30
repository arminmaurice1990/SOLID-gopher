package cmd

import (
	"cobra-cli-web-scraper/saver"
	"cobra-cli-web-scraper/scrape-manager"
	"cobra-cli-web-scraper/scraper"
	"github.com/spf13/cobra"
	"log"
)

var sav = saver.NewSaver()
var scrap = scraper.NewScraper()
var manager = scrape_manager.NewScrapeManager(scrap, sav)

var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape web pages into local files",
	Long:  "Scrape web pages into local files",
	Run: func(cmd *cobra.Command, args []string) {
		manager.Run(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
