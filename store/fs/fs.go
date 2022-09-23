package fs

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/4nth0/goaway/redirect"
)

type Client struct {
	Path  string
	store map[string]redirect.Redirect
}

func NewClient(path string) (*Client, error) {
	if path == "" {
		return nil, errors.New("EMPTY_PATH")
	}
	client := Client{
		Path:  path,
		store: make(map[string]redirect.Redirect),
	}

	if DatabaseFileExists(path) {
		err := client.Load()
		if err != nil {
			return nil, err
		}
	} else {
		err := client.Save()
		if err != nil {
			return nil, err
		}
	}

	return &client, nil
}

func DatabaseFileExists(path string) bool {
	_, err := ioutil.ReadFile(path)
	return err == nil
}

func (c *Client) Load() error {
	data, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &c.store)
}

func (c *Client) Store(id string, redirect redirect.Redirect) {
	c.store[id] = redirect
	c.Save()
}

func (c *Client) ByID(id string) (*redirect.Redirect, error) {
	if redirect, ok := c.store[id]; ok {
		return &redirect, nil
	}
	return nil, errors.New("redirect not found")
}

func (c *Client) RedirectExists(id string) bool {
	_, err := c.ByID(id)
	return err == nil
}

func (c *Client) Save() error {
	return ioutil.WriteFile(c.Path, []byte(c.Dump()), 0644)
}

func (c Client) Dump() string {
	b, _ := json.Marshal(c.store)
	return string(b)
}
