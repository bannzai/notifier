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
		// TODO: Add test cases.
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
