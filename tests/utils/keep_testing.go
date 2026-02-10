package utils

import (
	"context"
	"fmt"
	"time"
)

type keepTestingConfig struct {
	interval   time.Duration
	firstDelay time.Duration
}

type keepTestingOpt = func(*keepTestingConfig)

func WithTestingInterval(interval time.Duration) keepTestingOpt {
	return func(conf *keepTestingConfig) {
		conf.interval = interval
	}
}

func WithFirstTestDelay(delay time.Duration) keepTestingOpt {
	return func(conf *keepTestingConfig) {
		conf.firstDelay = delay
	}
}

func KeepTesting(testFn func() error, timeout time.Duration, opts ...keepTestingOpt) error {
	conf := keepTestingConfig{
		interval:   1 * time.Second,
		firstDelay: 0,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	timer := time.NewTimer(conf.firstDelay)
	defer timer.Stop()

	var lastErr error
	for {
		select {
		case <-ctx.Done():
			if err := testFn(); err == nil {
				return nil
			}
			return fmt.Errorf("keepTesting timed out after %v: %w", timeout, lastErr)

		case <-timer.C:
			err := testFn()
			if err == nil {
				return nil
			}

			lastErr = err
			timer.Reset(conf.interval)
		}
	}
}
