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

// fetchCOVID19Data ...
func fetchCOVID19Data(ctx context.Context, req *http.Request) <-chan LastValues {
	ch := make(chan LastValues)
	go func() {
		var r struct {
			Data LastValues `json:"data"`
		}
		body, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("fetchCOVID19Data: %v", err)
			return
		}
		defer body.Body.Close()
		err = json.NewDecoder(body.Body).Decode(&r)
		if err != nil {
			log.Printf("fetchCOVID19Data: %v", err)
			return
		}

		select {
		case ch <- LastValues{r.Data.Confirmed, r.Data.Deaths, r.Data.Recovered}:
		case <-ctx.Done():
		}
	}()
	return ch
}

func routine(sleep time.Duration) {
	cachedVal := LastValues{}
	const timeout = time.Second * 2
	for {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
		if err != nil {
			panic("internal error - misuse of NewRequestWithContext")
		}
		select {
		case newVal := <-fetchCOVID19Data(ctx, req):
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

func main() {
	log.SetPrefix(os.Args[0] + ": ")
	log.SetFlags(0)
	var timer time.Duration
	flag.DurationVar(&timer, "t", time.Hour, "interval between each api request")
	flag.Parse()
	routine(timer)
}
