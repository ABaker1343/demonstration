package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Result struct {
    Value int
    Err error
}

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			Res Result
		}{
			{0, Result{1, nil}},
			{1, Result{1, nil}},
			{2, Result{2, nil}},
			{3, Result{6, nil}},
		}

	for _, test := range tests {
		/* act */
        result := Result{}
        result.Value, result.Err = Factorial(test.n)
		/* assert */
		assert.Equal(test.Res.Value, result.Value)
        assert.Equal(test.Res.Err, result.Err)
	}
}
