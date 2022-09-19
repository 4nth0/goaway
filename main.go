package main

import (
	"log"
	"net/http"
	"strings"
)

type Condition struct {
	Source        string
	Operator      string
	ExpectedValue interface{}
}

type ConditionedRedirect struct {
	Value     string
	Condition Condition
}

type Redirect struct {
	Default    string
	Conditions []ConditionedRedirect
}

const SourceDelimiter = "."

func main() {
	redirects := map[string]Redirect{
		"1234": {
			Default: "https://google.com",
			Conditions: []ConditionedRedirect{
				{
					Value: "https://google.fr",
					Condition: Condition{
						Source:        "query.params.lang",
						Operator:      "eq",
						ExpectedValue: "fr",
					},
				},
				{
					Value: "https://google.it",
					Condition: Condition{
						Source:        "query.params.lang",
						Operator:      "eq",
						ExpectedValue: "it",
					},
				},
				{
					Value: "https://yahoo.com",
					Condition: Condition{
						Source:        "query.params.joke",
						Operator:      "eq",
						ExpectedValue: "1",
					},
				},
			},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if redirect, ok := redirects[strings.TrimPrefix(r.URL.Path, "/")]; ok {

			path := GetRedirectPath(redirect, r)

			w.Write([]byte(path))
			return
		}
		w.Write([]byte("Nop!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetRedirectPath(redirect Redirect, r *http.Request) string {
	for _, condition := range redirect.Conditions {
		splitedSource := strings.Split(condition.Condition.Source, SourceDelimiter)

		switch splitedSource[0] {
		case "query":
			if isValidQueryCondition(splitedSource[1], splitedSource[2], condition.Condition.ExpectedValue, r) {
				return condition.Value
			}
		}
	}
	return redirect.Default
}

func isValidQueryCondition(context, key string, expectedValue interface{}, r *http.Request) bool {
	switch context {
	case "params":
		return r.URL.Query().Get(key) == expectedValue
	}
	return false
}
