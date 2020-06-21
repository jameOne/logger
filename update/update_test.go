package update

import "testing"

func TestUpdates(t *testing.T) {

	var upd Logger

	upd = New("test update")

	exp := "test update"

	if upd.Update() != exp {
		t.Errorf(
			"&updateString %s != %s",
			upd,
			exp,
		)
	}
}
