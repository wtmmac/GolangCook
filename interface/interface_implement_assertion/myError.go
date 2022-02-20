package main

// MyError 统一错误处理，增加Code
type MyError struct {
	Message string
	Code    int
}

func NewMyError(message string, code int) error {
	return &MyError{Message: message, Code: code}
}

func (e *MyError) Error() string {
	return e.Message
}
