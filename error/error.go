package error

import "fmt"

// Logger variable defines itself as any variable that
// can describe itself as a string.
type Logger interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text.
func New(text string) Logger {
	return &errorString{text}
}

// Join returns an error that formats as the given text.
func Join(perr error, err error) Logger {
	return &errorString{fmt.Sprint(perr.Error(), ": ", err.Error())}
}
