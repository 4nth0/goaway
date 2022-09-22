package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/4nth0/goaway/redirect"
	"github.com/4nth0/goaway/rules"
)

func main() {
	client := redirect.NewClient()

	client.Register("1234", "https://default.com", []redirect.Conditions{
		{
			Rules: []rules.Condition{
				{
					Source:         "query.params.lang",
					Operator:       "eq",
					ValueToCompare: "fr",
				},
				{
					Source:         "time.hour",
					Operator:       "gt",
					ValueToCompare: 17,
				},
				{
					Source:         "time.hour",
					Operator:       "lt",
					ValueToCompare: 19,
				},
			},
			Value: "https://gotosleep.co",
		},
		{
			Rules: []rules.Condition{
				{
					Source:         "query.params.lang",
					Operator:       "eq",
					ValueToCompare: "fr",
				},
			},
			Value: "https://google.fr",
		},
		{
			Rules: []rules.Condition{
				{
					Source:         "query.params.lang",
					Operator:       "eq",
					ValueToCompare: "it",
				},
			},
			Value: "https://google.it",
		},
		{
			Rules: []rules.Condition{
				{
					Source:         "query.params.joke",
					Operator:       "eq",
					ValueToCompare: "1",
				},
			},
			Value: "https://yahoo.com",
		},
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/")

		if client.RedirectExists(id) {
			path := client.GetRedirectPath(id, r)

			w.Write([]byte(path))
			return
		}
		w.Write([]byte("Nop!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
