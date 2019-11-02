package parser

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewGitHub(t *testing.T) {
	tests := []struct {
		name string
		want GitHub
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitHub(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitHub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHub_Parse(t *testing.T) {
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name    string
		parser  GitHub
		args    args
		want    Content
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := GitHub{}
			got, err := parser.Parse(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitHub.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitHub.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
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
