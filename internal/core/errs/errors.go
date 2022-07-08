package errors

type NotNumberError struct {
	Message string
}

func (e *NotNumberError) Error() string {
	return e.Message
}
