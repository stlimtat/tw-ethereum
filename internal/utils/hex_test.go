package utils

import (
	"context"
	"testing"
)

func TestHexToInt64(t *testing.T) {
	tests := []struct {
		in   string
		want int64
	}{
		{"0xf", 15},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			ctx := context.Background()
			got := HexToInt64(ctx, tt.in)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
