package output

// Logger innterface defines itself as any variable that
// can describe itself as a string.
type Logger interface {
	Output() string
}

type outputString struct {
	s string
}

func (o *outputString) Output() string {
	return o.s
}

// New returns an output that formats as the given text.
func New(text string) Logger {
	return &outputString{text}
}
