package env

import "github.com/subosito/gotenv"

// LoadEnvFileError - Defines a struct for failed load env file
type LoadEnvFileError struct {
	msg string
}

// NewLoadEnvFileError - Returns a new instance of LoadEnvFileError
func NewLoadEnvFileError(msg string) *LoadEnvFileError {
	return &LoadEnvFileError{msg}
}

// Error - Implements error interface
func (l *LoadEnvFileError) Error() string {
	return l.msg
}

// LoadEnvFile - Tries to load an env file from a path given, if panicOnFailure == true panic istead returning error
func LoadEnvFile(path string, panicOnFailure bool) error {

	if err := gotenv.Load(path); err != nil {
		if panicOnFailure {
			panic(err)
		} else {
			return NewLoadEnvFileError("Failed to load env file")
		}
	}

	return nil
}
