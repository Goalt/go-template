package main

import (
	"testing"

	mock_main "github.com/Goalt/go-template/cmd/mocks"
	"github.com/golang/mock/gomock"
)

func Test_simpleFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_main.NewMocksimpleInterface(ctrl)
	m.EXPECT().SimpleInterfaceFunc(gomock.Eq(42)).Return("42")

	type args struct {
		i simpleInterface
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple",
			args: args{
				i: m,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simpleFunc(tt.args.i)
		})
	}
}
