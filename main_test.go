package main

import (
	"testing"
	"time"
)

func TestCompareInfos(t *testing.T) {
	lastValues := &LastValues{
		Confirmed: 1,
		Deaths:    0,
		Recovered: 0,
		UpdatedAt: time.Now(),
	}
	type args struct {
		currentValues LastValues
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Compare Infos returns true",
			args: args{currentValues: *lastValues},
			want: true,
		},
		{
			name: "Test Compare Infos returns false",
			args: args{currentValues: LastValues{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareInfos(tt.args.currentValues); got != tt.want {
				t.Errorf("CompareInfos() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
