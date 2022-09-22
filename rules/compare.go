package rules

import "golang.org/x/exp/constraints"

func compare[K constraints.Ordered](value K, operator string, expected K) bool {
	switch operator {
	case "gt":
		return value > expected
	case "lt":
		return value < expected
	case "eq":
		return value == expected
	case "gte":
		return value >= expected
	case "lte":
		return value <= expected
	case "neq":
		return value != expected
	}
	return false
}
