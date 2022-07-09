package utils

import "bytes"

func Separator(char string, lenght int) string {
	var buffer bytes.Buffer
	for i := 0; i < lenght; i++ {
		buffer.WriteString(char)
	}

	return buffer.String()
}
