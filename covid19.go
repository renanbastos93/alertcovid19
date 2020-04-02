package covid19

import (
	"context"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gen2brain/beeep"
)

// ICovid19 ...
type ICovid19 interface {
	Alert(title string, message string, icon string)
	String() string
}

// Covid19Info ...
type Covid19Info struct {
	Confirmed int `json:"confirmed,cases"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

// Covid19 ...
type Covid19 struct {
	ICovid19
	Covid19Info
}

func (c *Covid19Info) String() string {
	return fmt.Sprintf("Confirmed: %d, Deaths: %d, Recovered: %d", c.Confirmed, c.Deaths, c.Recovered)
}

// Alert ...
func (c *Covid19) Alert(title string, i string, icon string) {
	err := beeep.Alert(title, i, icon)
	if err != nil {
		fmt.Printf("rountine: %v", err)
	}
	return
}

func fetch(ctx context.Context, req *http.Request, ch chan Covid19Info) error {
	defer close(ch)
	var r Covid19Info
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
	case ch <- Covid19Info{r.Confirmed, r.Deaths, r.Recovered}:
	case <-ctx.Done():
	}
	return nil
}

// FetchCOVID19Data ...
func (c *Covid19) FetchCOVID19Data(ctx context.Context, url string) <-chan Covid19Info {
	ch := make(chan Covid19Info)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic("internal error - misuse of NewRequestWithContext")
	}
	go fetch(ctx, req, ch)
	return ch
}
