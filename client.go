package mailchimp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client manages communication with the Mailchimp API.
type Client struct {
	client  *http.Client
	baseURL *url.URL
	dc      string
	apiKey  string
}

// NewClient returns a new Mailchimp API client.  If a nil httpClient is
// provided, http.DefaultClient will be used. The apiKey must be in the format xyz-us11.
func NewClient(apiKey string, httpClient *http.Client) (ClientInterface, error) {
	if len(strings.Split(apiKey, "-")) != 2 {
		return nil, errors.New("Mailchimp API Key must be formatted like: xyz-zys")
	}
	dc := strings.Split(apiKey, "-")[1] // data center
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, err := url.Parse(fmt.Sprintf("https://%s.api.mailchimp.com/3.0", dc))
	if err != nil {
		return nil, err
	}
	return &Client{
		client:  httpClient,
		baseURL: baseURL,
		apiKey:  apiKey,
		dc:      dc,
	}, nil
}

// GetBaseURL ...
func (c *Client) GetBaseURL() *url.URL {
	return c.baseURL
}

// SetBaseURL ...
func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.baseURL = baseURL
}

func (c *Client) do(method string, path string, body interface{}) (*http.Response, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	apiURL := fmt.Sprintf("%s%s", c.GetBaseURL(), path)

	req, err := http.NewRequest(method, apiURL, buf)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth("", c.apiKey)

	return c.client.Do(req)
}

func (c *Client) doRequest(req MailchimpRequest) (*http.Response, error) {
	return c.do(
		req.Method,
		req.Path,
		req.Body,
	)
}

func (c *Client) ExecuteBatchOperation(b BatchOperation) (*http.Response, error) {
	return c.do("POST", "/batches", b.Requests)
}

func extractError(data []byte) (*ErrorResponse, error) {
	errorResponse := new(ErrorResponse)
	if err := json.Unmarshal(data, errorResponse); err != nil {
		return nil, err
	}
	return errorResponse, nil
}
