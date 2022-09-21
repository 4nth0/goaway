package rules

import (
	"net/http"
	"time"
)

func isValidTimeCondition(kind, operator string, expectedValue interface{}, r *http.Request) bool {
	t := time.Now()
	var valueToCompare int
	switch kind {
	case "hour":
		valueToCompare = t.Hour()
	}

	switch operator {
	case "gt":
		return valueToCompare > expectedValue.(int)
	case "lt":
		return valueToCompare < expectedValue.(int)
	case "eq":
		return valueToCompare == expectedValue.(int)
	case "gte":
		return valueToCompare >= expectedValue.(int)
	case "lte":
		return valueToCompare <= expectedValue.(int)
	case "neq":
		return valueToCompare != expectedValue.(int)
	}

	return false
}
