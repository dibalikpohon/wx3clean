package errors

import "fmt"

type What struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type DataNotExistsError struct {
	Message string `json:"message"`
	What    What   `json:"what"`
}

func NewDataNotExistsError(field, value string) *DataNotExistsError {
	what := What{
		Field: field,
		Value: value,
	}

	return &DataNotExistsError{
		Message: "Data tidak ditemukan",
		What:    what,
	}
}

func (e *DataNotExistsError) Error() string {
	return fmt.Sprintf("%v. (%v: %v)",
		e.Message, e.What.Field, e.What.Value)
}
