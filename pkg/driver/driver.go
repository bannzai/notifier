package driver

import "net/http"

type Driver interface {
	Drive(r *http.Request) error
}
