package rules

import "net/http"

func isValidQueryCondition(context, key, operator string, expectedValue interface{}, r *http.Request) bool {
	valueToCompare := ""

	switch context {
	case "params":
		valueToCompare = r.URL.Query().Get(key)
	}

	return compare(valueToCompare, operator, expectedValue.(string))
}
