package httpClient

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	http.Client
	// baseUrl is the root URL for all invocations of the client.
	baseUrl *url.URL
	// for cli, we set keepAlive to false so that the connection
	// will be closed after each request
	keepAlive bool
}

func NewHTTPClient(config *Config) (*Client, error) {
	baseUrl, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid base url: %s\n", config.BaseUrl))
	}
	return &Client{
		Client: http.Client{
			Timeout: defaultTimeout,
		},
		baseUrl:   baseUrl,
		keepAlive: config.KeepAlive,
	}, nil
}

func (c *Client) Close() {
	c.Client.CloseIdleConnections()
}

func (c *Client) Send(req *http.Request) ([]byte, int, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, -1, errors.Wrapf(err, "failed to send request to %s\n", req.URL.String())
	}
	defer resp.Body.Close()

	code := resp.StatusCode

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, code, errors.Wrapf(err, "failed to read response body from %s\n", req.URL.String())
	}

	return response, code, nil
}

func (c *Client) verb(verb string) *Request {
	return newRequest(c, verb)
}

func (c *Client) Get() *Request {
	return c.verb("GET")
}

func (c *Client) Post() *Request {
	return c.verb("POST")
}

func (c *Client) Put() *Request {
	return c.verb("PUT")
}

func (c *Client) Delete() *Request {
	return c.verb("DELETE")
}

func (c *Client) Patch() *Request {
	return c.verb("PATCH")
}

// getUrlFromConfig parse and create a base url from the config
func getUrlFromConfig(config *Config) (*url.URL, error) {
	return url.Parse(config.BaseUrl)
}
