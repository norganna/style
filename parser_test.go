package style

import "testing"

func Test_FindSequence(t *testing.T) {
	tests := [][5]string{
		{"‹", "›", " ‹LI› ", "LI", ""},
		{"‹", "›", "‹fn:body›", "fn", "body"},
		{"‹", "›", "‹x:‹fn:body››", "fn", "body"},
		{"A", "B", "Afn:bodyB", "fn", "body"},
		{"{{", "}}", "{{fn:x{}x}}", "fn", "x{}x"},
		{":", ":", ":fn:x:", "fn", ""},  // This case is non-obvious, don't be stupid and use : as both start and end.
		{":", "::", ":fn:x::", "x", ""}, // Again, don't do this, it matches :x::.
		{"::", "::", "::fn:x::", "fn", "x"},
	}

	for _, test := range tests {
		_, _, fn, body := findSequence([]rune(test[0]), []rune(test[1]), []rune(test[2]))
		if fn != test[3] {
			t.Errorf(
				"Failed matching “%s”, expected function “%s” got “%s”",
				test[2], test[3], fn,
			)
		}
		if body != test[4] {
			t.Errorf(
				"Failed matching “%s”, expected body “%s” got “%s”",
				test[2], test[4], body,
			)
		}
	}
}
