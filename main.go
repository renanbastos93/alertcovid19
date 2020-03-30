package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	timer = flag.Duration("timer", time.Hour, "This parameter define time to verify API ")
)

// LastValues ...
type LastValues struct {
	Confirmed uint `json:"confirmed"`
	Deaths    uint `json:"deaths"`
	Recovered uint `json:"recovered"`
	UpdatedAt time.Time
}

type response struct {
	Data LastValues `json:"data"`
}

var lastValues LastValues

// Exported ...
const (
	URL string = "https://covid19-brazil-api.now.sh/api/report/v1/brazil"
	IMG string = "https://static.poder360.com.br/2020/03/23312-868x644.png"
)

func (l *LastValues) String() string {
	return fmt.Sprintf("Confirmed: %d, Deaths: %d, Recovered: %d", l.Confirmed, l.Deaths, l.Recovered)
}

// GetDataCOVID19 ...
func GetDataCOVID19(ctx context.Context) (*LastValues, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	errCh := make(chan error, 1)
	respCh := make(chan LastValues, 1)
	go func() {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errCh <- err
		}
		var resp *response
		err = json.Unmarshal(body, &resp)
		if err != nil {
			errCh <- err
		}
		respCh <- LastValues{
			Confirmed: resp.Data.Confirmed,
			Deaths:    resp.Data.Deaths,
			Recovered: resp.Data.Recovered,
			UpdatedAt: time.Now(),
		}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err = <-errCh:
		return nil, err
	case resp := <-respCh:
		return &resp, nil
	}
}

// CompareInfos ...
func CompareInfos(currentValues LastValues) bool {
	if lastValues.Confirmed != currentValues.Confirmed ||
		lastValues.Deaths != currentValues.Deaths ||
		lastValues.Recovered != currentValues.Recovered {
		return true
	}
	return false
}

func routine(duration time.Duration) {
	const timeout = time.Second * 2
	for ctx, cancel := context.WithTimeout(context.Background(), timeout); ; cancel() {
		currentValue, err := GetDataCOVID19(ctx)
		if err != nil {
			panic(err)
		}

		if CompareInfos(*currentValue) {
			err := beeep.Alert("COVID-19 Brazil", currentValue.String(), IMG)
			if err != nil {
				panic(err)
			}
			lastValues = *currentValue
		}
		fmt.Println("opa time? ", duration, *timer)
		time.Sleep(duration)
	}
}

func main() {
	flag.Parse()
	routine(time.Duration(*timer))
}
