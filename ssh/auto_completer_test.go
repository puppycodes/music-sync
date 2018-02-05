package ssh

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type autoCompleterTestCase struct {
	line    string
	newLine []string
}

var autoCompleterTestCases = []autoCompleterTestCase{
	{
		line:    "c",
		newLine: []string{"lear "},
	},
	{
		line:    "e",
		newLine: []string{"xit "},
	},
	{
		line:    "",
		newLine: []string{"clear ", "exit ", "help ", "ls "},
	},
	{
		line:    "l",
		newLine: []string{"s "},
	},
}

func TestSshAutoCompleter_Do(t *testing.T) {
	autoCompleter := sshAutoCompleter{}

	for _, c := range autoCompleterTestCases {
		nl, _ := autoCompleter.Do([]rune(c.line), len(c.line))
		resultNL := make([]string, len(nl))
		for i, l := range nl {
			resultNL[i] = string(l)
		}
		assertStringSliceEqual(t, c.newLine, resultNL, "autoCompleter returned")
	}
}

func assertStringSliceEqual(t *testing.T, expected, actual []string, name string) {
	if assert.Equal(t, len(expected), len(actual), "%s wrong number of strings (%v, %v)", name, expected, actual) {
		sort.Strings(expected)
		sort.Strings(actual)
		for i := range expected {
			assert.Equal(t, expected[i], actual[i], "%s wrong string at index %d (%v vs %s)", name, i, expected, actual)
		}
	}
}