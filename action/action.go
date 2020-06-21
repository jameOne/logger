// Package action is one component of a complete
// logging utility intended to promote the writing of
// intuitive logs.
package action

// Logger interface defines itself as any structure
// that implements the Action() method.
type Logger interface {
	Action() string
}

// Unexported struct type used to implement the
// Action interface.
type actionString struct {
	s string
}

// This method ensures actionString will meet
// requirements of the Action interface.
func (a *actionString) Action() string {
	return a.s
}

// New returns an action that formats as the given text.
func New(text string) Logger {
	return &actionString{text}
}
