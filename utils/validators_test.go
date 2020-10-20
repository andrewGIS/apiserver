package utils

import (
	"errors"
	"testing"
)

func TestTable(t *testing.T) {
	cases := []struct {
		Email  string
		Result bool
		Err    error
	}{
		{
			Email:  "andrew@gmail.com",
			Result: true,
			Err:    nil,
		},
		{
			Email:  "andrew@gmail.",
			Result: false,
			Err:    nil,
		},
		{
			Email:  "",
			Result: false,
			Err:    errors.New("Empty email"),
		},
		{
			Email:  "andrew.gmail.com",
			Result: false,
			Err:    errors.New("Invalid email"),
		},
	}

	for _, testCase := range cases {
		res := IsEmailValid(testCase.Email)
		if res != testCase.Result {
			t.Errorf("Excepted %v, got is %v", testCase.Result, res)
		}
	}
}
