package input

import "testing"

func TestIOs(t *testing.T) {

	var in Logger

	in = New("test input")

	exp := "test input"

	if in.IO() != exp {
		t.Errorf(
			"&inputString %s != %s",
			in,
			exp,
		)
	}
}
