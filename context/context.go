package context


import "time"

// START OMIT
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
	// Done returns a channel that is closed when this Context is canceled
	// or times out.
	Done() <-chan struct{}

	// Err indicates why this context was canceled, after the Done channel
	// is closed.
	Err() error

	// Deadline returns the time when this Context will be canceled, if any.
	Deadline() (deadline time.Time, ok bool)

	// Value returns the value associated with key or nil if none.
	Value(key interface{}) interface{}
}
// END OMIT
