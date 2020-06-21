package input

// Logger variable defines itself as any variable that
// can describe itself as a string.
type Logger interface {
	IO() string
}

type inputString struct {
	s string
}

func (i *inputString) IO() string {
	return i.s
}

// New returns an input that formats as the given text.
func New(text string) Logger {
	return &inputString{text}
}
