package message

import (
	"testing"
)

func TestParamsWithNoParams(t *testing.T) {
	params := Params([]string{})
	if params.String() != "" {
		t.Fail()
	}
}

func TestParamsWithOneParam(t *testing.T) {
	const p = "one"

	params := Params([]string{p})
	if params.String() != " " + p {
		t.Fail()
	}
}

func TestParamsWithFourteenParams(t *testing.T) {
	params := Params([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"})

	if params.String() != " 1 2 3 4 5 6 7 8 9 10 11 12 13 14" {
		t.Error(params.String())
	}
}

func TestParamsWithFourteenPlusTrailingParams(t *testing.T) {
	params := ParamsT([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14"}, "I am trailing")

	if params.String() != " 1 2 3 4 5 6 7 8 9 10 11 12 13 14 :I am trailing" {
		t.Error(params.String())
	}
}

func TestParamsWithFifteenParams(t *testing.T) {
	params := Params([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"})

	if params.String() != " 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15" {
		t.Error(params.String())
	}
}
