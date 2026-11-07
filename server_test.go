package main

import "testing"

func TestCountHealthyServers(t *testing.T) {
	tests := []struct {
		name    string
		results []bool
		want    int
	}{
		{
			name:    "mixed servers",
			results: []bool{true, false, true, false, true},
			want:    3,
		},
		{
			name:    "empty server list",
			results: []bool{},
			want:    0,
		},
		{
			name:    "all unhealthy",
			results: []bool{false, false, false},
			want:    0,
		},
		{
			name:    "all healthy",
			results: []bool{true, true, true},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountHealthyServers(tt.results)

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
