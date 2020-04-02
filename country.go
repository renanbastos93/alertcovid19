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
func (c *Country) GetData(country string, cache Covid19) {
	url := urlCountry + country
	timeout := 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	select {
	case newVal := <-c.Covid19.FetchCOVID19Data(ctx, url):
		// if Covid19 != newVal {
		log.Println("cache L- ", cache)
		c.Covid19.Alert("COVID-19 "+country, newVal.String(), "")
		// cachedVal = newVal
		// }
	case <-ctx.Done():
		log.Printf("rountine: %v", ctx.Err())
	}
	cancel()
	return
}
