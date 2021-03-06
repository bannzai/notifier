package parser

import (
	"reflect"
	"testing"

	"github.com/bannzai/notifier/pkg/testutil"
)

func TestGitHub_parseBody(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		g       GitHub
		args    args
		want    Content
		wantErr bool
	}{
		{
			name: "parse successfully when read truth json for mention on comment pattern",
			g:    NewGitHub(),
			args: args{
				body: testutil.ReadFile(t, testutil.CallerDirectoryPath(t)+"/testdata/github_comment.json"),
			},
			want: Content{
				LinkURL:     "https://github.com/bannzai/notifier/pull/1#issuecomment-549011949",
				UserNames:   []string{"bannzai"},
				ContentType: GitHubMentionContent,
			},
		},
		{
			name: "parse successfully when read truth json for assigned to pull-request pattern",
			g:    NewGitHub(),
			args: args{
				body: testutil.ReadFile(t, testutil.CallerDirectoryPath(t)+"/testdata/github_assigned_pull_request.json"),
			},
			want: Content{
				LinkURL:     "https://github.com/bannzai/notifier/pull/1",
				UserNames:   []string{"bannzai"},
				ContentType: GitHubAssignedContent,
			},
		},
		{
			name: "parse successfully when read truth json for request reviewed to pull-request pattern",
			g:    NewGitHub(),
			args: args{
				body: testutil.ReadFile(t, testutil.CallerDirectoryPath(t)+"/testdata/github_request_reviewed_pull_request.json"),
			},
			want: Content{
				LinkURL:     "https://github.com/bannzai/notifier/pull/1",
				UserNames:   []string{"bannzai"},
				ContentType: GitHubRequestReviewedContent,
			},
		},
		{
			name: "parse successfully when read truth json for single request reviewed to pull-request pattern",
			g:    NewGitHub(),
			args: args{
				body: testutil.ReadFile(t, testutil.CallerDirectoryPath(t)+"/testdata/github_request_reviewed_pull_request_has_single_request_reviewer.json"),
			},
			want: Content{
				LinkURL:     "https://github.com/bannzai/notifier/pull/1",
				UserNames:   []string{"bannzai"},
				ContentType: GitHubRequestReviewedContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GitHub{}
			got, err := g.parseBody(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitHub.parseBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitHub.parseBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
