package httpClient

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Request struct {
	client *Client

	verb    string
	path    string
	params  url.Values
	headers http.Header

	body io.Reader
}

func newRequest(client *Client, verb string) *Request {
	return &Request{
		client: client,
		verb:   verb,
	}
}

// SetResource sets the request resource path.
func (r *Request) SetResource(name ...string) *Request {
	r.path = strings.Join(name, "/")
	return r
}

// SetURL overwrites existing path and parameters with the given url.
func (r *Request) SetURL(url string) *Request {
	r.client.baseUrl = nil
	r.path = url
	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

// Param creates a query parameter with the given string value.
func (r *Request) Param(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

func (r *Request) Params(paramName string, values []string) *Request {
	for _, val := range values {
		r.Param(paramName, val)
	}
	return r
}

func (r *Request) Body(body []byte) *Request {
	r.body = bytes.NewBuffer(body)
	return r
}

func (r *Request) URL() (string, error) {
	var (
		final *url.URL
		err   error
	)

	if r.client.baseUrl != nil {
		final = r.client.baseUrl
		final.Path = path.Join(final.Path, r.path)
	} else {
		final, err = url.Parse(r.path)
		if err != nil {
			return "", errors.Wrap(err, "failed to parse url")
		}
	}

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	final.RawQuery = query.Encode()
	return final.String(), nil
}

func (r *Request) BuildRequest(ctx context.Context) (*http.Request, error) {
	url, err := r.URL()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.verb, url, r.body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.headers
	return req, nil
}

func (r *Request) Send(ctx context.Context) ([]byte, error) {
	req, err := r.BuildRequest(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Send(req)
}
