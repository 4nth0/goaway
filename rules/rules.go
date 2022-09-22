package rules

import (
	"net/http"
	"strings"
)

type Condition struct {
	Source         string      `json:"source"`
	Operator       string      `json:"operator"`
	ValueToCompare interface{} `json:"valueToCompare"`
}

const SourceDelimiter = "."

func IsConditionSucceeding(rules []Condition, r *http.Request) bool {
	for _, rule := range rules {
		splitedSource := strings.Split(rule.Source, SourceDelimiter)

		switch splitedSource[0] {
		case "query":
			if !isValidQueryCondition(splitedSource[1], splitedSource[2], rule.Operator, rule.ValueToCompare, r) {
				return false
			}
		case "time":
			if !isValidTimeCondition(splitedSource[1], rule.Operator, rule.ValueToCompare, r) {
				return false
			}
		}
	}
	return true
}
