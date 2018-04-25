package model

type CustomError struct {
	Message string
	StatusCode  int
}

func (c CustomError) Error() string {
	return c.Message
}

func (c CustomError) Status() int {
	return c.StatusCode
}
