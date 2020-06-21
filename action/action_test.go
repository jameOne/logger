package action

import "testing"

func TestActions(t *testing.T) {

	var act Logger

	act = New("test action")

	exp := "test action"

	if act.Action() != exp {
		t.Errorf(
			"&actionString %s != %s",
			act,
			exp,
		)
	}
}
