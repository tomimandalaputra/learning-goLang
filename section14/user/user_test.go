package user

import "testing"

func TestCheckUsername(t *testing.T) {
	got := CheckUsername("testing")
	if got != true {
		t.Errorf("got %v\nwant %v", got, true)
	}
}

func TestCheckUsernameTable(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "to short", input: "test1", want: false},
		{name: "empty", input: "", want: false},
		{name: "contains admin", input: "adminuser", want: false},
		{name: "valid", input: "greateusername", want: true},
	}

	for _, tc := range testCases {
		got := CheckUsername(tc.input)
		if got != tc.want {
			t.Errorf("got %v\nwant %v", got, tc.want)
		}
	}
}
func TestCheckUsernameTableWithSubtest(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "to short", input: "test", want: false},
		{name: "empty", input: "", want: false},
		{name: "contains admin", input: "adminuser", want: false},
		{name: "valid", input: "greateusername", want: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUsername(tc.input)
			if got != tc.want {
				t.Fatalf("got %v\nwant %v", got, tc.want)
			}
		})
	}
}
