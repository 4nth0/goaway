package stats

import (
	"context"
	"encoding/json"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type Hit struct {
	ID      string
	Method  string              `json:"method"`
	Headers map[string][]string `json:"headers"`
	URL     *url.URL            `json:"url"`
	Body    string              `json:"body,omitempty"`
}

type StatsWriter interface {
	WriteLine(string) error
	Close()
}

type Client struct {
	Hits   chan Hit
	writer StatsWriter
}

func NewClient(writer StatsWriter) *Client {
	return &Client{
		Hits:   make(chan Hit),
		writer: writer,
	}
}

func (c *Client) Collect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			c.writer.Close()
			close(c.Hits)
			return
		case hit := <-c.Hits:
			line, _ := json.Marshal(hit)
			err := c.writer.WriteLine(string(line))
			if err != nil {
				log.Error(err)
			}
		}
	}
}
