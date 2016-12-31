// Package accesslog provider a baa middleware for log http access.
package accesslog

import (
	"time"

	"gopkg.in/baa.v1"
)

// Accesslog returns a baa middleware for log http access
func Accesslog() baa.HandlerFunc {
	return func(c *baa.Context) {
		start := time.Now()

		c.Next()

		c.Baa().Logger().Printf("%s %s %s %v %v", c.RemoteAddr(), c.Req.Method, c.URL(false),  c.Resp.Status(), time.Since(start))
	}
}
