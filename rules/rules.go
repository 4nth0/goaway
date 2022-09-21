package rules

import (
	"net/http"
	"strings"
)

type Condition struct {
	Source        string
	Operator      string
	ExpectedValue interface{}
}

type ConditionedRedirect struct {
	Value string
	Rules []Condition
}

const SourceDelimiter = "."

func IsConditionSucceeding(rules []Condition, r *http.Request) bool {
	for _, rule := range rules {
		splitedSource := strings.Split(rule.Source, SourceDelimiter)

		switch splitedSource[0] {
		case "query":
			if !isValidQueryCondition(splitedSource[1], splitedSource[2], rule.ExpectedValue, r) {
				return false
			}
		case "time":
			if !isValidTimeCondition(splitedSource[1], rule.Operator, rule.ExpectedValue, r) {
				return false
			}
		}
	}
	return true
}
