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
	"github.com/renanbastos93/crawler-covid19-rs"
)

// Exported ...
const (
	IMG string = "https://static.poder360.com.br/2020/03/23312-868x644.png"
	URL        = "https://covid19-brazil-api.now.sh/api/report/v1/brazil"
)

// LastValues ...
type LastValues struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

func (l LastValues) String() string {
	return fmt.Sprintf("Confirmed: %d, Deaths: %d, Recovered: %d", l.Confirmed, l.Deaths, l.Recovered)
}

// fetch runs on its own goroutine
func fetch(ctx context.Context, req *http.Request, ch chan LastValues) error {
	defer close(ch)
	var r struct {
		Data LastValues `json:"data"`
	}
	body, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("fetchCOVID19Data: %v", err)
		return err
	}
	defer body.Body.Close()
	err = json.NewDecoder(body.Body).Decode(&r)
	if err != nil {
		log.Printf("fetchCOVID19Data: %v", err)
		return err
	}

	select {
	case ch <- LastValues{r.Data.Confirmed, r.Data.Deaths, r.Data.Recovered}:
	case <-ctx.Done():
	}
	return nil
}

// fetchCOVID19Data ...
func fetchCOVID19Data(ctx context.Context) <-chan LastValues {
	ch := make(chan LastValues)
	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		panic("internal error - misuse of NewRequestWithContext")
	}
	go fetch(ctx, req, ch)
	return ch
}

func routine(sleep time.Duration) {
	cachedVal := LastValues{}
	const timeout = time.Second * 2
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		select {
		case newVal := <-fetchCOVID19Data(ctx):
			if cachedVal != newVal {
				err := beeep.Alert("COVID-19 Brazil", newVal.String(), IMG)
				if err != nil {
					log.Printf("rountine: %v", err)
				}
				cachedVal = newVal
			}
		case <-ctx.Done():
			log.Printf("rountine: %v", ctx.Err())
		}
		cancel()
		log.Printf("sleeping for %s", sleep)
		time.Sleep(sleep)
	}
}

func getStateRSByCrawler() {
	body := crawler.GetData()
	cities, values := crawler.GetCityValue(body)
	cities, values = crawler.ReplaceValue(cities, values)
	list := crawler.SetCitiesValues(cities, values)
	fmt.Printf("%+v :: xx", list)
}

func flags(timer *time.Duration, state *string, city *string) {
	flag.DurationVar(timer, "t", time.Hour, "interval between each api request")
	flag.StringVar(state, "state", "-", "Chose your filter state")
	flag.StringVar(city, "city", "-", "Chose your filter city")
	flag.Parse()
}

type myFlags interface{}

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	var (
		timer time.Duration
		state string
		city  string
	)
	flags(&timer, &state, &city)
	routine(timer)
	// getStateRSByCrawler()
}
