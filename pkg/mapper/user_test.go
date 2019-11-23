package mapper

import (
	"os"
	"reflect"
	"testing"

	"github.com/bannzai/notifier/pkg/testutil"
)

func Test_fetchUsers(t *testing.T) {
	tests := []struct {
		before  func()
		after   func()
		name    string
		want    []User
		wantErr bool
	}{
		{
			before: func() {
				path := testutil.CallerDirectoryPath(t) + "/testdata/test.yaml"
				if err := os.Setenv("YAML_FILE_PATH", path); err != nil {
					t.Errorf("can not load YAML_FILE_PATH from %s, got error of %v", path, err)
				}
			},
			after: func() {
				if err := os.Unsetenv("YAML_FILE_PATH"); err != nil {
					t.Errorf("YAML_FILE_PATH unset env failure got error of %v", err)
				}
			},
			name: "successfully fetch users from test.yml",
			want: []User{
				{
					ID: "bannzai",
					GitHub: GitHub{
						Login: "kojiki",
					},
					Slack: Slack{
						ID: "BIK0NY93C",
					},
				},
				{
					ID: "yudai.hirose",
					GitHub: GitHub{
						Login: "kingkong999",
					},
					Slack: Slack{
						ID: "xyz",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()
			got, err := fetchUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRemoteURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "path begin http://",
			args: args{
				path: "http://example.com",
			},
			want: true,
		},
		{
			name: "path begin https://",
			args: args{
				path: "https://example.com",
			},
			want: true,
		},
		{
			name: "path is absolute local file path",
			args: args{
				path: "/Users/bannzai/notifier/test.yml",
			},
			want: false,
		},
		{
			name: "path is relative local file path",
			args: args{
				path: "test.yml",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRemoteURL(tt.args.path); got != tt.want {
				t.Errorf("isRemoteURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
