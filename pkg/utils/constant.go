package util

import "fmt"

type StatusTicket string

const (
	Sold      = "Sold"
	Available = "Available"
)

func (status StatusTicket) Validate() error {
	switch status {
	case Sold, Available:
		return nil
	}

	return fmt.Errorf("Invalid status ticket")
}
