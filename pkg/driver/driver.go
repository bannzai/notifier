package driver

import "net/http"

type Driver interface {
	Drive(r *http.Request) error
}

func Drive(r *http.Request, drivers ...Driver) error {
	var driver Drivers = drivers
	return driver.Drive(r)
}
