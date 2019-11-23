package mapper

import (
	"reflect"
	"testing"

	"github.com/bannzai/notifier/pkg/sender"
)

func Test_extractUserFromGitHub(t *testing.T) {
	type args struct {
		users                []User
		githubUserName       string
		extractedContentType sender.ContentType
	}
	tests := []struct {
		name  string
		args  args
		want  Slack
		want1 bool
	}{
		{
			name: "extract slack user",
			args: args{
				users: []User{
					{
						ID: "bannzai",
						GitHub: GitHub{
							Login: "github",
						},
						Slack: Slack{
							ID: "slack",
						},
					},
					{
						ID: "yudai.hirose",
						GitHub: GitHub{
							Login: "xyz",
						},
						Slack: Slack{
							ID: "abc",
						},
					},
				},
				githubUserName:       "github",
				extractedContentType: sender.SlackContentType,
			},
			want: Slack{
				ID: "slack",
			},
			want1: true,
		},
		{
			name: "no matched github user name",
			args: args{
				users: []User{
					{
						ID: "bannzai",
						GitHub: GitHub{
							Login: "github",
						},
						Slack: Slack{
							ID: "slack",
						},
					},
					{
						ID: "yudai.hirose",
						GitHub: GitHub{
							Login: "xyz",
						},
						Slack: Slack{
							ID: "abc",
						},
					},
				},
				githubUserName:       "hogehoge",
				extractedContentType: sender.SlackContentType,
			},
			want:  Slack{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := extractUserFromGitHub(tt.args.users, tt.args.githubUserName, tt.args.extractedContentType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractUserFromGitHub() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("extractUserFromGitHub() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
