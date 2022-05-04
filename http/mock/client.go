package mock

import (
	"fmt"
	"net/http"

	httpclient "github.com/HybriStratus/test-github-groups/http"
)

// NewMockClient is a wrapper for http.Client
func NewMockClient() httpclient.Client {
	return Client{}
}

// HTTPClient is an interface that wraps the Do function from http
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a struct that holds a list of Responses
type Client struct {
	Responses map[string]map[string][]http.Response
}

// SetResponses is a method that add to the list of responses held in Client
func (c *Client) SetResponses(method string, url string, response http.Response) {
	if len(c.Responses) == 0 {
		c.Responses = make(map[string]map[string][]http.Response)
	}

	if len(c.Responses[url][method]) == 0 {
		if c.Responses[url] == nil {
			c.Responses[url] = make(map[string][]http.Response)
		}
		if c.Responses[url][method] == nil {
			c.Responses[url][method] = []http.Response{}
		}
		c.Responses[url][method] = append(c.Responses[url][method], response)
		return
	}
	c.Responses[url][method] = append(c.Responses[url][method], response)
}

// Do overrides the http Do method for the mock client to use it
func (c Client) Do(req *http.Request) (*http.Response, error) {
	if responses, ok := c.Responses[req.URL.String()][req.Method]; ok {
		response := responses[0]
		c.Responses[req.URL.String()][req.Method] = responses[1:]
		return &response, nil
	}
	return nil, fmt.Errorf("no response set")
}
