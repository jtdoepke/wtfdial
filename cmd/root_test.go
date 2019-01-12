package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWTFLevelArg_bad_input(t *testing.T) {
	result, err := ParseWTFLevelArg("not a number")
	assert.Error(t, err)
	assert.Equal(t, 0.0, result)
}

func ExampleParseWTFLevelArg_int() {
	lvl, err := ParseWTFLevelArg("50")
	if err != nil {
		panic(err)
	}
	fmt.Println(lvl)
	// Output: 0.5
}
func ExampleParseWTFLevelArg_percent() {
	lvl, err := ParseWTFLevelArg("50%")
	if err != nil {
		panic(err)
	}
	fmt.Println(lvl)
	// Output: 0.5
}

func ExampleParseWTFLevelArg_division() {
	lvl, err := ParseWTFLevelArg("50/100")
	if err != nil {
		panic(err)
	}
	fmt.Println(lvl)
	// Output: 0.5
}

func ExampleParseWTFLevelArg_float() {
	lvl, err := ParseWTFLevelArg("0.5")
	if err != nil {
		panic(err)
	}
	fmt.Println(lvl)
	// Output: 0.5
}
