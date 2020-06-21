package yield

import "testing"

func TestYields(t *testing.T) {

	var yld Logger

	yld = New("test yield")

	exp := "test yield"

	if yld.Yield() != exp {
		t.Errorf(
			"&errorString %s != %s",
			yld,
			exp,
		)
	}
}
