package wtf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWTFService(t *testing.T) {
	wtf := New()

	err := wtf.SetLevel(SetLevelRequest{Username: "matt", Level: 0.11}, nil)
	assert.NoError(t, err)

	err = wtf.SetLevel(SetLevelRequest{Username: "jaye", Level: 0.01}, nil)
	assert.NoError(t, err)

	err = wtf.SetLevel(SetLevelRequest{Username: "liza", Level: 0.3}, nil)
	assert.NoError(t, err)

	var result float64
	err = wtf.Avg(nil, &result)
	assert.NoError(t, err)

	assert.InDelta(t, 0.14, result, 0.001)
}
