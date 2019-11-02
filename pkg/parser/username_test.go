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
		{
			name: "Not contains user names",
			args: args{
				target: "Not contains user names text",
			},
			want: []string{},
		},
		{
			name: "contains one user name",
			args: args{
				target: "@bannzai is user name",
			},
			want: []string{"bannzai"},
		},
		{
			name: "contains some user name",
			args: args{
				target: "@bannzai and @yudai-hirose is user names",
			},
			want: []string{"bannzai", "yudai-hirose"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userNames(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
