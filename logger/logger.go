package logger

// Logger is a simple interface which is used to log critical messages (aka panics).
type Logger func(errorMessage string)
