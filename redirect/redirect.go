package redirect

import (
	"encoding/json"
	"errors"
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

type Client struct {
	store map[string]Redirect
}

func NewClient() *Client {
	return &Client{
		store: make(map[string]Redirect),
	}
}

func (c *Client) Register(id string, defaultValue string, redirects []Conditions) Redirect {
	r := Redirect{
		Default:   defaultValue,
		Redirects: redirects,
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
	for _, condition := range redirect.Redirects {
		if rules.IsConditionSucceeding(condition.Rules, r) {
			return condition.Value
		}
	}
	return redirect.Default
}

func (c Client) Dump() string {
	b, _ := json.MarshalIndent(c.store, "", "	")
	return string(b)
}
