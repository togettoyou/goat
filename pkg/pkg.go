package pkg

import (
	"goat-layout/pkg/conf"
	"goat-layout/pkg/log"
)

// Reset all
func Reset() error {
	if err := conf.Reset(); err != nil {
		return err
	}
	log.Reset(conf.Log.Level)
	return nil
}
