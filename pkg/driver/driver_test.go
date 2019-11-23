package driver

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		parser Parser
		sender Sender
	}
	tests := []struct {
		name string
		args args
		want Driver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.parser, tt.args.sender); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_Drive(t *testing.T) {
	type fields struct {
		Parser Parser
		Sender Sender
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := Driver{
				Parser: tt.fields.Parser,
				Sender: tt.fields.Sender,
			}
			if err := driver.Drive(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Driver.Drive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
