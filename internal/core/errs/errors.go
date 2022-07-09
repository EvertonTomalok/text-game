package errs

type NotNumberError struct {
	Message string
}

func (e *NotNumberError) Error() string {
	return e.Message
}

type InvalidRoundsNumber struct {
	Message string
}

func (e *InvalidRoundsNumber) Error() string {
	return e.Message
}
