package main

/* Item
TODO: Rename to ItemDB */
type Item struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	URL         string `db:"url"`
	ImageID     string `db:"image_id"`
	Person      string `db:"person"`
	Quantity    int    `db:"quantity"`
	Created     string `db:"created"`
	Deleted     bool   `db:"deleted"`
	Completed   bool   `db:"completed"`
	CompletedOn string `db:"completed_on"`
}

/* ItemJSON */
type ItemJSON struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	ImageID     string `json:"image_id"`
	Person      string `json:"person"`
	Quantity    int    `json:"quantity"`
	Created     string `json:"created"`
	Deleted     bool   `json:"deleted"`
	Completed   bool   `json:"completed"`
	CompletedOn string `json:"completed_on"`
}

/* Response
TODO: Split Response struct into Response and ErrorResponse */
type Response struct {
	Response string `json:"response"`
	Errno    int    `json:"errno"`
	Error    string `json:"error"`
}
