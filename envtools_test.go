package envtools

import "testing"

func TestMustGetenv(t *testing.T) {
	val := "test_env_var_val"
	t.Setenv("ENVTOOLS_TEST_VAR", val)
	if got := MustGetenv("ENVTOOLS_TEST_VAR"); got != val {
		t.Fatalf("got %q, want %q", got, val)
	}
}

func TestMustGetenvPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustGetenv did not panic, want panic")
		}
	}()

	MustGetenv("ENVTOOLS_TEST_VAR")
}

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
