package envtools

import "testing"

func TestIsTruthy(t *testing.T) {
	cases := []struct {
		name string
		val  string
		want bool
	}{
		{"empty", "", false},
		{"positive_number", "1", true},
		{"zero", "0", false},
		{"negative_number", "-1", false},
		{"true_string", "on", true},
		{"true_string_capitalized", "ENABLED", true},
		{"false_string", "off", false},
		{"false_string_capitalized", "NO", false},
		{"unknown_nonempty_string", "foo", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("ENVTOOLS_TEST_VAR", tc.val)
			if got := IsTruthy("ENVTOOLS_TEST_VAR"); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsTruthyNotSet(t *testing.T) {
	want := false
	if got := IsTruthy("ENVTOOLS_TEST_VAR"); got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
