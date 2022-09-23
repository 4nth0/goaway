package redirect

import (
	"net/http"

	"github.com/4nth0/goaway/rules"
)

type Conditions struct {
	Value string            `json:"value"`
	Rules []rules.Condition `json:"rules"`
}

type Redirect struct {
	Default   string       `json:"default"`
	Redirects []Conditions `json:"redirects"`
}

type DBClient interface {
	ByID(id string) (*Redirect, error)
	Store(id string, redirect Redirect)
}

type Client struct {
	db DBClient
}

func NewClient(db DBClient) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) Register(id string, defaultValue string, redirects []Conditions) Redirect {
	r := Redirect{
		Default:   defaultValue,
		Redirects: redirects,
	}
	c.db.Store(id, r)
	return r
}

func (c *Client) GetRedirectPath(id string, r *http.Request) string {
	redirect, err := c.db.ByID(id)
	if err != nil {
		return ""
	}
	for _, condition := range redirect.Redirects {
		if rules.IsConditionSucceeding(condition.Rules, r) {
			return condition.Value
		}
	}
	return redirect.Default
}
