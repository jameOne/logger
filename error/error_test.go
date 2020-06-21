package error

import "testing"

func TestErrors(t *testing.T) {

	err := New("test error")

	exp := "test error"

	if err.Error() != exp {
		t.Errorf(
			"&errorString %s != %s",
			err,
			exp,
		)
	}

	perr := New("previous test error")
	err = Join(perr, err)

	exp = "previous test error: test error"

	if err.Error() != exp {
		t.Errorf(
			"&errorString %s != %s",
			err,
			exp,
		)
	}

}
