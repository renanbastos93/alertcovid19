package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	timer = flag.Int64("timer", 3600, "This parameter define time in seconds to verify API ")
)

func init() {
	flag.Parse()
}

// LastValues ...
type LastValues struct {
	Confirmed uint `json"confirmed"`
	Deaths    uint `json"deaths"`
	Recovered uint `json"recovered"`
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
func GetDataCOVID19() (currentValue LastValues, err error) {
	res, err := http.Get(URL)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var resp *response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	currentValue = LastValues{
		Confirmed: resp.Data.Confirmed,
		Deaths:    resp.Data.Deaths,
		Recovered: resp.Data.Recovered,
		UpdatedAt: time.Now(),
	}
	return
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
	for {
		currentValue, err := GetDataCOVID19()
		if err != nil {
			panic(err)
		}

		if CompareInfos(currentValue) {
			err := beeep.Alert("COVID-19 Brazil", currentValue.String(), IMG)
			if err != nil {
				panic(err)
			}
			lastValues = currentValue
		}
		fmt.Println("opa time? ", duration, *timer)
		time.Sleep(duration * time.Second)
	}
}

func main() {
	lastValues = LastValues{}
	routine(time.Duration(*timer))
}
