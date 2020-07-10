package main

/*
 * Item
 */
type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Person   string `json:"person"`
	Quantity int    `json:"quantity"`
}

/*
 * Response
 */
type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
