package client

// Client represents a knowstellation client
// that is able to consume knowstellationd API
type Client struct {
	endpoint string
}

// NewClient creates a knowstellation client
func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

// Execute the client logic
func (*Client) Execute() (bool, error) {
	return true, nil
}
