package flexy

type Client struct {
	Address string `json:"address"`
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
	}
}

func (c *Client) Query() (q *Query) {
	return &Query{
		QAddress: c.Address,
	}
}
