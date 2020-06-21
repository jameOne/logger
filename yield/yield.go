package yield

// Logger interface defines itself as any variable that
// can describe itself as a string.
type Logger interface {
	Yield() string
}

type yieldString struct {
	s string
}

// Yeilds returns a yield that formats as the given text.
func (y *yieldString) Yield() string {
	return y.s
}

// New returns an update that formats as the given text.
func New(text string) Logger {
	return &yieldString{text}
}
