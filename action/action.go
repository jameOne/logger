// Package action is one component of a complete
// logging utilty intended to promiote the writing of
// intuitive logs.
package action

// Logger interface defines itself as any structure
// that implements the Action() method.
type Logger interface {
	Action() string
}

// Unexported stuct type used to implement the
// Action interface.
type actionString struct {
	s string
}

// This method ensures actionString will meet
// requirments of the Action interface.
func (a *actionString) Action() string {
	return a.s
}

// New returns an action that formats as the given text.
func New(text string) Logger {
	return &actionString{text}
}
