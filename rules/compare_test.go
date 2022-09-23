package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Arg interface {
	int | int64 | float64 | string
}

func TestCompare(t *testing.T) {
	assert.True(t, compare("a", "eq", "a"))
	assert.False(t, compare("a", "eq", "b"))

	assert.True(t, compare("c", "gt", "a"))
	assert.False(t, compare("c", "gt", "d"))

	assert.True(t, compare("a", "lt", "c"))
	assert.False(t, compare("d", "lt", "c"))

	assert.True(t, compare("b", "gte", "b"))
	assert.True(t, compare("b", "gte", "a"))
	assert.False(t, compare("a", "gte", "b"))

	assert.True(t, compare("a", "lte", "a"))
	assert.True(t, compare("a", "lte", "b"))
	assert.False(t, compare("b", "lte", "a"))

	assert.True(t, compare("a", "neq", "b"))
	assert.False(t, compare("a", "neq", "a"))

	assert.True(t, compare(1, "eq", 1))
	assert.False(t, compare(1, "eq", 2))

	assert.True(t, compare(1, "gt", 0))
	assert.False(t, compare(1, "gt", 2))

	assert.True(t, compare(1, "lt", 2))
	assert.False(t, compare(2, "lt", 1))

	assert.True(t, compare(1, "gte", 1))
	assert.True(t, compare(1, "gte", 0))
	assert.False(t, compare(1, "gte", 2))

	assert.True(t, compare(1, "lte", 1))
	assert.True(t, compare(1, "lte", 2))
	assert.False(t, compare(2, "lte", 1))

	assert.True(t, compare(1, "neq", 2))
	assert.False(t, compare(1, "neq", 1))

	assert.False(t, compare(1, "invalid", 2))
}
