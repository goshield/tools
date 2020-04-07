package tools

// PanicOnError panics if input value is not nil
func PanicOnError(err interface{}) {
	if err != nil {
		panic(err)
	}
}

// Must panics if there is an error
func Must(errors ...interface{}) {
	for _, err := range errors {
		PanicOnError(err)
	}
}
