// Package apod provides utilities for fetching APOD photos from NASA's API.
package apod

type APOD struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	Copyright   string `json:"copyright",omitempty`
	Date        Date   `json:"date"`
	URL         string `json:"url"`
	HDUrl       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
}

type Client struct {
	apikey string
}

func New(apikey string) *Client {
	var c = &Client{}
	c.SetAPIKey(apikey)
	return c
}

func (c *Client) SetAPIKey(apikey string) {
	c.apikey = apikey
}

func (c *Client) Latest() (*APOD, error) {
	params := c.params()
	return get(params)
}

func (c *Client) ForDate(date Date) (*APOD, error) {
	params := c.params()
	params["date"] = date.String()
	return get(params)
}

func (c *Client) params() map[string]string {
	return map[string]string{
		"api_key": c.apikey,
	}
}
