package redirect

import (
	"errors"
	"net/http"

	"github.com/4nth0/goaway/rules"
)

type ConditionedRedirect struct {
	Value string
	Rules []rules.Condition
}

type Redirect struct {
	Default    string
	Conditions []ConditionedRedirect
}

type Client struct {
	store map[string]Redirect
}

func NewClient() *Client {
	return &Client{
		store: make(map[string]Redirect),
	}
}

func (c *Client) Register(id string, defaultValue string, redirects []ConditionedRedirect) Redirect {
	r := Redirect{
		Default:    defaultValue,
		Conditions: redirects,
	}
	c.store[id] = r
	return r
}

func (c *Client) ByID(id string) (*Redirect, error) {
	if redirect, ok := c.store[id]; ok {
		return &redirect, nil
	}
	return nil, errors.New("redirect not found")
}

func (c *Client) RedirectExists(id string) bool {
	_, err := c.ByID(id)
	return err == nil
}

func (c *Client) GetRedirectPath(id string, r *http.Request) string {
	redirect, err := c.ByID(id)
	if err != nil {
		return ""
	}
	for _, condition := range redirect.Conditions {
		if rules.IsConditionSucceeding(condition.Rules, r) {
			return condition.Value
		}
	}
	return redirect.Default
}
