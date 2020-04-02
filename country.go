package covid19

import (
	"context"
	"log"
	"time"
)

// Country ...
type Country struct {
	Covid19
}

var urlCountry = "https://corona.lmao.ninja/countries/"

// GetData ...
func (c *Country) GetData(country string, cache *Country) {
	url := urlCountry + country
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	select {
	case newVal := <-c.Covid19.FetchCOVID19Data(ctx, url):
		if cache.Covid19.Covid19Info != newVal {
			c.Covid19.Alert("COVID-19 "+country, newVal.String(), "")
		}
	case <-ctx.Done():
		log.Printf("rountine: %v", ctx.Err())
	}
	cancel()
	return
}
