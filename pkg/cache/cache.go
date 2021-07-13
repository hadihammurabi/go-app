package cache

import "time"

type Cache interface {
	Set(string, interface{}, ...time.Duration) error
	Get(string) (interface{}, error)
	IsAvailable() error
}
