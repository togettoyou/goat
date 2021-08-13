package pkg

import (
	"goat/pkg/conf"
	"goat/pkg/log"
)

// Reset all
func Reset() error {
	if err := conf.Reset(); err != nil {
		return err
	}
	log.Reset()
	return nil
}
