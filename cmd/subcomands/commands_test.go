package subcomands

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestAdd(t *testing.T) {
	type args struct {
		command *cli.Command
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.command)
		})
	}
}
