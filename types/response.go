package types

// define entity

type Pagination struct {
	current  int
	pageSize int
	total    int
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
