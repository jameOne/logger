package output

import "testing"

func TestOutputs(t *testing.T) {

	var out Logger

	out = New("test output")

	exp := "test output"

	if out.Output() != exp {
		t.Errorf(
			"&outputString %s != %s",
			out,
			exp,
		)
	}
}
