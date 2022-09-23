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

	return compare(valueToCompare, operator, int(expectedValue.(float64)))
}
