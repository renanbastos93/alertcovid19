package main

import (
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
				Confirmed: 100,
				Deaths:    7,
				Recovered: 1,
			},
			want: "Confirmed: 100, Deaths: 7, Recovered: 1",
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
