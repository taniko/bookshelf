package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDevelopment(t *testing.T) {
	tests := []struct {
		env  string
		want bool
	}{
		{
			env:  "development",
			want: true,
		}, {
			env:  "production",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			t.Setenv("APP_ENV", tt.env)
			assert.Equal(t, tt.want, IsDevelopment())
		})
	}
}

func TestIsProduction(t *testing.T) {
	tests := []struct {
		env  string
		want bool
	}{
		{
			env:  "development",
			want: false,
		}, {
			env:  "production",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			t.Setenv("APP_ENV", tt.env)
			assert.Equal(t, tt.want, IsProduction())
		})
	}
}
