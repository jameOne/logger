package update

// Logger variable defines itself as any variable that
// can describe itself as a string.
type Logger interface {
	Update() string
}

type updateString struct {
	s string
}

func (u *updateString) Update() string {
	return u.s
}

// New returns an update that formats as the given text.
func New(text string) Logger {
	return &updateString{text}
}
