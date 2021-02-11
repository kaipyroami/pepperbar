/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/gorilla/feeds"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
	"io/ioutil"
)

// twitterCmd represents the twitter command
var twitterCmd = &cobra.Command{
	Use:   "twitter",
	Short: "Pull Twitter data and publish to an RSS",
	Long: `This module will pull a specified number of tweets 
	from a user and publish it to an RSS, Atom or JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting twitter posts from " + username)

		runScraper(username, count)
	},
}

// Variables for twitter flag functionality
var username, outfile, feedtype string
var count int

func init() {
	rootCmd.AddCommand(twitterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twitterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twitterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// TODO: Add optional feed title, link, description and author flags.
	twitterCmd.PersistentFlags().StringVarP(&username,"username", "u", "twitter", "Twitter username")
	twitterCmd.PersistentFlags().IntVarP(&count, "count", "c", 25, "Help message for toggle")
	twitterCmd.PersistentFlags().StringVarP(&outfile,"outfile", "f", "feed.xml", "Name of output file.")
	twitterCmd.PersistentFlags().StringVarP(&feedtype,"format", "t", "rss", "Format for the feed generated. Atom, RSS or JSON")
	twitterCmd.MarkFlagRequired("username")
}


func runScraper(twituser string, numtweets int){
	scraper := twitterscraper.New()


	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Feed generated with pepperbar",
		Link:        &feeds.Link{Href: "https://github.com/kaipyroami/pepperbar"},
		Description: "A command line web scraper to RSS utility.",
		Author:      &feeds.Author{Name: "Kyle Crockett", Email: "kyle@kyle-crockett.com"},
		Created:     now,
	}


	for tweet := range scraper.GetTweets(context.Background(), twituser, numtweets) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		newItem := feeds.Item{
			Id: tweet.ID,
			Link:        &feeds.Link{Href: tweet.PermanentURL},
			Description: tweet.Text,
			Author:      &feeds.Author{Name: tweet.Username},
			Created:     time.Unix(tweet.Timestamp,0),
		}

		feed.Items = append(feed.Items, &newItem)
	}

	var feedOut string
	var err error

	switch feedtype {
	case "rss":
		feedOut, err = feed.ToRss()
		if err != nil {
			log.Fatal(err)
		}
	case "atom":
		feedOut, err = feed.ToAtom()
		if err != nil {
			log.Fatal(err)
		}
	case "json":
		feedOut, err = feed.ToJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
	// TODO: Need to fix this filepath issue.
	modeVal, _ := strconv.ParseUint("0775", 8, 32)
	fm := os.FileMode(modeVal)
	err = ioutil.WriteFile(outfile, []byte(feedOut), fm)
	if err != nil {
		panic(err)
	}


}
