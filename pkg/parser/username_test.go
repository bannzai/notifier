package parser

import (
	"reflect"
	"testing"
)

func Test_userNames(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userNames(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
