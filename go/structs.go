package main

/*
 Item
*/
// To be renamed to ItemDB
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
 ItemJSON
*/
type ItemJSON struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Person   string `json:"person"`
	Quantity int    `json:"quantity"`
	Deleted  bool   `json:"deleted"`
}

/*
 Response
*/
type Response struct {
	Response string `json:"response"`
	Errno    int    `json:"errno"`
	Error    string `json:"error"`
}
