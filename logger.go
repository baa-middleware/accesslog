// Package logger provider a baa middleware for log http access.
package logger

import (
	"time"

	"github.com/go-baa/baa"
)

// Logger returns a baa middleware for log http access
func Logger() baa.MiddlewareFunc {
	return func(h baa.HandlerFunc) baa.HandlerFunc {
		return func(c *baa.Context) error {
			start := time.Now()

			if err := h(c); err != nil {
				c.Error(err)
			}

			c.Baa().Logger().Printf("log %s %s %v", c.Req.Method, c.Req.RequestURI, time.Since(start))

			return nil
		}
	}
}
