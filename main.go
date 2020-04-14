package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

// Exported ...
const (
	myIP        = "http://ip-api.com/json"
	IMG         = "https://static.poder360.com.br/2020/03/23312-868x644.png"
	URL         = "https://corona.lmao.ninja/countries/"
	StateBrazil = "https://covid19-brazil-api.now.sh/api/report/v1/brazil/uf/"
	TIMEOUT     = time.Second * 2
)

type geoIP struct {
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

// CovidStatus has the number of cases, deaths and recovered patients
type CovidStatus struct {
	Cases     int `json:"cases"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

func (s CovidStatus) String() string {
	message := "Cases: %d, Deaths: %d, Recovered: %d"
	return fmt.Sprintf(message, s.Cases, s.Deaths, s.Recovered)
}

// fetch runs on its own goroutine
func fetch(ctx context.Context, req *http.Request, ch chan CovidStatus) error {
	defer close(ch)
	var s CovidStatus
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("fetchCOVID19Data: %v", err)
		return err
	}
	defer body.Body.Close()
	err = json.NewDecoder(body.Body).Decode(&s)
	if err != nil {
		log.Printf("fetchCOVID19Data: %v", err)
		return err
	}
	select {
	case ch <- CovidStatus{s.Cases, s.Deaths, s.Recovered}:
	case <-ctx.Done():
	}
	return nil
}

func fetchCovidStatus(ctx context.Context, country string) <-chan CovidStatus {
	ch := make(chan CovidStatus)
	url := URL + country
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic("internal error - misuse of NewRequestWithContext")
	}
	go fetch(ctx, req, ch)
	return ch
}

func routine(sleep time.Duration, country string) {
	lastStatus := CovidStatus{}
	for {
		ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
		select {
		case newStatus := <-fetchCovidStatus(ctx, country):
			if newStatus != lastStatus {
				err := beeep.Alert("COVID-19 "+country, newStatus.String(), IMG)
				if err != nil {
					log.Printf("routine: %v", err)
				}
				lastStatus = newStatus
			}
		case <-ctx.Done():
			log.Printf("routine: %v", ctx.Err())
		}
		cancel()
		time.Sleep(sleep)
	}
}

func getCountryByGeoIP() geoIP {
	client := http.Client{Timeout: TIMEOUT}
	resp, err := client.Post(
		myIP,
		"application/json; charset=utf8",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var ip geoIP
	err = json.NewDecoder(resp.Body).Decode(&ip)
	if err != nil {
		log.Fatalf("Oops, we cannot get your location, please verify your network.")
	}
	return ip
}

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	var timer time.Duration
	flag.DurationVar(&timer, "t", time.Hour, "interval between each api request")
	flag.Parse()
	ip := getCountryByGeoIP()
	routine(timer, ip.CountryCode)
}
