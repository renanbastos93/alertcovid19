package main

import (
	"testing"
)

func TestLastValues_String(t *testing.T) {
	type fields struct {
		Confirmed uint
		Deaths    uint
		Recovered uint
	}
	tests := struct {
		name   string
		fields fields
		want   string
	}{
		name: "Test format message to show in the notification",
		fields: fields{
			Confirmed: 100,
			Deaths:    7,
			Recovered: 1,
		},
		want: "Confirmed: 100, Deaths: 7, Recovered: 1",
	}
	t.Run(tests.name, func(t *testing.T) {
		l := &LastValues{
			Confirmed: tests.fields.Confirmed,
			Deaths:    tests.fields.Deaths,
			Recovered: tests.fields.Recovered,
		}
		if got := l.String(); got != tests.want {
			t.Errorf("LastValues.String() = %v, want %v", got, tests.want)
		}
	})
}
