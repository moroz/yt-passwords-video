package types

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult []ValidationError

func (r *ValidationResult) Add(field, message string) {
	*r = append(*r, ValidationError{Field: field, Message: message})
}

func (r ValidationResult) Valid() bool {
	return len(r) == 0
}
