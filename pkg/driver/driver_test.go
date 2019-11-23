package driver

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bannzai/notifier/pkg/parser"
	gomock "github.com/golang/mock/gomock"
)

func TestDriver_Drive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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
		{
			name: "successfully drive",
			fields: fields{
				Parser: func() Parser {
					mock := NewMockParser(ctrl)
					mock.EXPECT().Parse(gomock.Any()).Return(parser.Content{LinkURL: "https://notifier.example.com"}, nil)
					return mock
				}(),
				Sender: func() Sender {
					mock := NewMockSender(ctrl)
					mock.EXPECT().Send(parser.Content{LinkURL: "https://notifier.example.com"}).Return(nil)
					return mock
				}(),
			},
			args: args{
				r: &http.Request{},
			},
			wantErr: false,
		},
		{
			name: "when parser is failed",
			fields: fields{
				Parser: func() Parser {
					mock := NewMockParser(ctrl)
					mock.EXPECT().Parse(gomock.Any()).Return(parser.Content{}, errors.New(""))
					return mock
				}(),
				Sender: func() Sender {
					mock := NewMockSender(ctrl)
					return mock
				}(),
			},
			args: args{
				r: &http.Request{},
			},
			wantErr: true,
		},
		{
			name: "when sender is failed",
			fields: fields{
				Parser: func() Parser {
					mock := NewMockParser(ctrl)
					mock.EXPECT().Parse(gomock.Any()).Return(parser.Content{LinkURL: "https://notifier.example.com"}, nil)
					return mock
				}(),
				Sender: func() Sender {
					mock := NewMockSender(ctrl)
					mock.EXPECT().Send(parser.Content{LinkURL: "https://notifier.example.com"}).Return(errors.New(""))
					return mock
				}(),
			},
			args: args{
				r: &http.Request{},
			},
			wantErr: true,
		},
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
