package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// LastValues ...
type LastValues struct {
	Cases     int `json:"cases"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

func (l LastValues) String() string {
	message := "Cases: %d, Deaths: %d, Recovered: %d"
	return fmt.Sprintf(message, l.Cases, l.Deaths, l.Recovered)
}

// fetch runs on its own goroutine
func fetch(ctx context.Context, req *http.Request, ch chan LastValues) error {
	defer close(ch)
	var r LastValues
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

	bodyBytes, _ := ioutil.ReadAll(body.Body)
	log.Println("WOOW: :: ", string(bodyBytes))

	select {
	case ch <- LastValues{r.Cases, r.Deaths, r.Recovered}:
	case <-ctx.Done():
	}
	return nil
}

// fetchCOVID19Data ...
func fetchCOVID19Data(ctx context.Context, country string) <-chan LastValues {
	ch := make(chan LastValues)
	url := URL + country
	log.Println("URL ::: ", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic("internal error - misuse of NewRequestWithContext")
	}
	go fetch(ctx, req, ch)
	return ch
}

func routine(sleep time.Duration, country string) {
	cachedVal := LastValues{}
	for {
		ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
		select {
		case newVal := <-fetchCOVID19Data(ctx, country):
			log.Println("newval :: ", newVal)
			if cachedVal != newVal {
				err := beeep.Alert("COVID-19 "+country, newVal.String(), IMG)
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
	log.Println("IP :: ", ip)
	return ip
}

func main() {
	// log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	var timer time.Duration
	flag.DurationVar(&timer, "t", time.Hour, "interval between each api request")
	flag.Parse()
	ip := getCountryByGeoIP()
	routine(timer, ip.CountryCode)
}
