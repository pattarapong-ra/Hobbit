package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "3252353253252g",
			want:"24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
			assert.Equal(t,tt.want,tt.name)
		})
	}
}
