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

type InvalidQuestion struct {
	Message string
}

func (e *InvalidQuestion) Error() string {
	return e.Message
}

type InvalidComplementaryQuestion struct {
	Message string
}

func (e *InvalidComplementaryQuestion) Error() string {
	return e.Message
}
