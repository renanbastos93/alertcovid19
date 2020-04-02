package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/renanbastos93/alertcovid19"
)

func filter(country, state, city string) (covid struct{}, f string) {
	if city != "-" {
		return
	}
	if state != "-" {
		return
	}
	if country != "-" {
		x := new(covid19.Country)
		log.Printf("\n\nopa :: %+v\n\n", x)
		x.GetData(country, x)
		return
	}
	return
}

func routine(sleep time.Duration, covid struct{}, ff string) {
	// var cachedVal struct{}
	for {
		// covid19.GetData(ff, cachedVal)
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
	filter(country, state, city)
	// routine(timer, f, ff)
	// getStateRSByCrawler()
}
