// Package logger provider a baa middleware for log http access.
package logger

import (
	"time"

	"github.com/go-baa/baa"
)

// Logger returns a baa middleware for log http access
func Logger() baa.HandlerFunc {
	return func(c *baa.Context) {
		start := time.Now()

		c.Next()

		c.Baa().Logger().Printf("%s %s %v %v", c.Req.Method, c.Req.RequestURI, c.Resp.Status(), time.Since(start))
	}
}
