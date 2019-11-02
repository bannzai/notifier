package driver

import (
	"fmt"
	"net/http"
)

type Drivers []Driver

func (drivers Drivers) Drive(r *http.Request) error {
	for _, driver := range drivers {
		if err := driver.Drive(r); err != nil {
			return fmt.Errorf("driver.Drive() error %w", err)
		}
	}
	return nil
}
