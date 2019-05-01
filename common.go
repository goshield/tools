package tools

// Panic panics if input value is not nil
func Panic(err interface{}) {
	if err != nil {
		panic(err)
	}
}

// Must panics if there is an error
func Must(errors ...interface{}) {
	for _, err := range errors {
		Panic(err)
	}
}