// Package accesslog provider a baa middleware for log http access.
package accesslog

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-baa/baa"
)

// Logger returns a baa middleware for log http access
func Logger() baa.HandlerFunc {
	return func(c *baa.Context) {
		start := time.Now()

		c.Next()

		c.Baa().Logger().Printf("%s %s %s %v %v\n", c.RemoteAddr(), c.Req.Method, c.URL(true), c.Resp.Status(), time.Since(start))
	}
}

// LoggerWithWriter returns a baa middleware with writer for log http access
func LoggerWithWriter(f io.Writer) baa.HandlerFunc {
	return func(c *baa.Context) {
		start := time.Now()

		c.Next()

		fmt.Fprintf(f, "%s %s %s %v %v\n", c.RemoteAddr(), c.Req.Method, c.URL(true), c.Resp.Status(), time.Since(start))
	}
}

// LoggerWithFile returns a baa middleware with file write for log http access
func LoggerWithFile(file string) baa.HandlerFunc {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("accesslog.LoggerWithFile: open logfile [%s] err: %v", file, err)
	}
	return LoggerWithWriter(f)
}
