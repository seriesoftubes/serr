// Package serr is the simplest way I could think of to return errors with a human readable stacktrace.
// I'm not saying it's fast or that you should use it everywhere,
// but if you just want a stacktrace with your error,
// `serr.New` is a drop-in replacement for `fmt.Errorf` and `errors.New`.
//
// Usage:
//   err := serr.New("yeah, it failed here because %v", badStuff)
//	 // <meanwhile, back in main()...>
//	 if err := reallyDeepFunctionCall(); err != nil {
//       fmt.Println("got err:", err.Error(), "because\n" + err.Stack())
//   }
package serr

import (
	"fmt"
	"runtime/debug"
)

// A Serr is a stacktrace combined with an error.
type Serr struct{ msg, stacktrace string }

// Error satisfies the `error` type interface.
func (s *Serr) Error() string { return s.msg }

// Stack returns the stacktrace captured when the Serr was formed.
func (s *Serr) Stack() string { return s.stacktrace }

// New returns a Serr with a stacktrace captured for the current goroutine.
func New(x string, args ...interface{}) error {
	msg := fmt.Sprintf(x, args...)
	stacktrace := fmt.Sprintf("%s", debug.Stack())
	return &Serr{msg, stacktrace}
}
