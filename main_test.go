package flago

import (
	"errors"
	"reflect"
	"testing"
)

type TestScenario struct {
	Description string
	Input       []string
	Expected    map[string]string
}

var ErrNotFound = errors.New("not found")

var testCases = []TestScenario{
	{"Simple", []string{"go_file_name", "alice", "1", "bob", "2", "charlie", "3"}, map[string]string{"alice": "1", "bob": "2", "charlie": "3"}},
	{"Incomplete", []string{"go_file_name", "alice", "1", "bob", "2", "charlie"}, map[string]string{"alice": "1", "bob": "2", "charlie": ""}},
	{"Remove Dash", []string{"go_file_name", "-alice", "1", "bob", "2", "-charlie"}, map[string]string{"alice": "1", "bob": "2", "charlie": ""}},
}

func TestGetArgs(t *testing.T) {
	for _, testCase := range testCases {

		t.Run(testCase.Description, func(t *testing.T) {
			actual, err := getPurgedArgs(testCase.Input)

			if err != nil {
				t.Errorf("Arguments parsing failed: %v\n", err)
			}

			equals := reflect.DeepEqual(testCase.Expected, actual)

			if !equals {
				t.Error("Output comparison failed")
			}
		})
	}
}
