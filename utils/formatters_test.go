package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatAlias(t *testing.T) {
	var expectedAlias = "b_twist"

	assert.Equal(t, expectedAlias, FormatAlias("B-Twist"))
	assert.Equal(t, expectedAlias, FormatAlias("B Twist"))
	assert.Equal(t, expectedAlias, FormatAlias("b_Twist"))
	assert.Equal(t, expectedAlias, FormatAlias("b twist"))
}
