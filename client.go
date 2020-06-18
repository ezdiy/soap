package soap

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type Client struct {
	*http.Client
	URL string
	Schema string
}

func NewClient(url, schema string) *Client {
	return &Client{
		Client: &http.Client{},
		URL: url,
		Schema: schema,
	}
}

const header = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"

// Encode a SOAP request into XML body
func (c *Client) Encode(name string, args Values) (rd io.Reader, err error) {
	x := &Body{}
	x.Body.Item.XMLName = xml.Name{
		Local: name,
		Space: c.Schema,
	}
	x.Marshal(args)
	b, err := xml.MarshalIndent(x, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(append([]byte(header), b...)), nil
}

// Decode a soap response from XML body
func (c *Client) Decode(r io.Reader) (v Values, err error) {
	x := &Body{}
	err = xml.NewDecoder(r).Decode(x)
	if err != nil {
		return
	}
	return x.Unmarshal(), nil
}

// Perform a single RPC call (encode request, http, decode response)
func (c *Client) Call(name string, args Values) (Values, error) {
	body, err := c.Encode(name, args)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", c.URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return c.Decode(resp.Body)
}
