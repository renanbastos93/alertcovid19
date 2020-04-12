package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLastValues_String(t *testing.T) {
	tests := []struct {
		name   string
		fields LastValues
		want   string
	}{
		{
			name: "Test format message to show in the notification",
			fields: LastValues{
				Cases:     100,
				Deaths:    7,
				Recovered: 1,
			},
			want: "Cases: 100, Deaths: 7, Recovered: 1",
		},
	}
	for _, v := range tests {
		v := v
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			if got := v.fields.String(); got != v.want {
				t.Errorf("LastValues.String() = %v, want %v", got, v.want)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  bool
	}{
		{"ok json", strings.NewReader(`{"data": {"cases": 10, "deaths": 10, "recovered": 10}}`), true},
		{"bad json", strings.NewReader(`{"data": "cases": 10, "deaths": 10, "recovered": 10}}`), false},
	}
	for _, v := range tests {
		v := v
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				io.Copy(w, v.input)
			}))
			defer ts.Close()
			req, err := http.NewRequest("GET", ts.URL, nil)
			if err != nil {
				panic("misuse of NewRequest")
			}
			ch := make(chan LastValues)
			go func() { // drain the fetch output out
				_ = <-ch
			}()
			if err := fetch(context.TODO(), req, ch); (err == nil) != v.want {
				t.Errorf("fetch: expected: %v got: %v", v.want, err)
			}
		})
	}
}
