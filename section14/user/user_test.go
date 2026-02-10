package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUsername(t *testing.T) {
	assert.True(t, CheckUsername("test123"))
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
		assert.Equal(t, tc.want, got, "got %v\nwant %v", got, tc.want)
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
			assert.Equal(t, tc.want, got, "got %v\nwant %v", got, tc.want)
		})
	}
}

func TestLoginSuccess(t *testing.T) {
	err, ok := Login("testusername")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestLoginFailure(t *testing.T) {
	err, ok := Login("name")
	assert.Error(t, err)
	assert.False(t, ok)
}
