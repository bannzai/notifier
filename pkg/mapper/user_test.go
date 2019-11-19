package mapper

import (
	"reflect"
	"testing"
)

func Test_fetchUsers(t *testing.T) {
	tests := []struct {
		name    string
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
