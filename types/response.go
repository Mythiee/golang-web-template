package types

// define entity

type Pagination struct {
	current  int
	pageSize int
	total    int
}

type Response struct {
	code    int
	message string
	data    interface{}
}
