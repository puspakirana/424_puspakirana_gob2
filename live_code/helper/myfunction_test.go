package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailMyFunction(t *testing.T) {
	result := myFunction()

	require.Equal(t, "0 2 4 8 10 ", result, "Result is false")
	fmt.Println("TestFailFunction Eksekusi Terhenti")
}

func TestMyFunction(t *testing.T) {
	result := myFunction()

	assert.Equal(t, "0 3 6 9 ", result, "Result is true")
	fmt.Println("TestFunction Eksekusi Terhenti")

}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  string
		expected string
		errMsg   string
	}{
		{
			request:  myFunction(),
			expected: "0 3 6 9 ",
			errMsg:   "Result is true",
		},
		{
			request:  myFunction(),
			expected: "0 2 4 6 8 ",
			errMsg:   "Result is false",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
