package rules

import "net/http"

func isValidQueryCondition(context, key string, expectedValue interface{}, r *http.Request) bool {
	switch context {
	case "params":
		return r.URL.Query().Get(key) == expectedValue
	}
	return false
}
