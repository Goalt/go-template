package main

import (
	"testing"

	"github.com/Goalt/go-template/cmd/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_someFunction(t *testing.T) {
	type args struct {
		intf Interf
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mIn := mocks.Interf{}
			mIn.On("Test", mock.AnythingOfType("int")).Return(func(i int) string {
				return "123"
			})

			someFunction(&mIn)
		})
	}
}
