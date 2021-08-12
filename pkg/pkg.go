package pkg

import (
	"goat/pkg/conf"
)

// Reset all
func Reset() error {
	if err := conf.Reset(); err != nil {
		return err
	}
	return nil
}
