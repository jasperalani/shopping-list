package main

/*
 * Item
 */
//type Item struct {
//	ID       int    `json:"id"`
//	Name     string `json:"name"`
//	URL      string `json:"url"`
//	ImageURL string `json:"image_url"`
//	Person   string `json:"person"`
//	Quantity int    `json:"quantity"`
//	Deleted  bool   `json:"deleted"`
//}
type Item struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	URL      string `db:"url"`
	ImageURL string `db:"image_url"`
	Person   string `db:"person"`
	Quantity int    `db:"quantity"`
	Deleted  bool   `db:"deleted"`
}

/*
 * Response
 */
type Response struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
