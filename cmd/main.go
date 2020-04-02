package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/renanbastos93/alertcovid19"
	"github.com/renanbastos93/crawler-covid19-rs"
)

func filter(country, state, city string) (interface{}, string) {
	if city != "-" {
		return
	}
	if state != "-" {
		return
	}
	if country != "-" {
		return covid19.Country, country
	}
}

func routine(sleep time.Duration, f interface{}, ff string) {
	var cachedVal struct{}
	for {
		f.Covid19.GetData(ff, cachedVal)
		log.Printf("sleeping for %s", sleep)
		time.Sleep(sleep)
	}
}

func flags(timer *time.Duration, country *string, city *string, state *string) {
	flag.DurationVar(timer, "t", time.Hour, "interval between each api request")
	flag.StringVar(country, "country", "-", "Chose your filter country")
	flag.StringVar(state, "state", "-", "Chose your filter state")
	flag.StringVar(city, "city", "-", "Chose your filter city")
	flag.Parse()
}

type myFlags interface{}

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	var (
		timer   time.Duration
		country string
		state   string
		city    string
	)
	flags(&timer, &country, &state, &city)
	f, ff := filter(country, state, city)
	routine(timer, f, ff)
	// getStateRSByCrawler()
}
