## Things ToDo

[ ] TODO

-Log functoin : `logs.LogIt(enum, funcName, errMsg, error)`

- **Enums**
	```go
	const (
		LogFatal = iota
		LogWarn  = iota
		LogInfo  = iota
	)
	```

- **Errors**
	```
		ErrUnAuthorized   = fmt.Errorf("unauthorized")
		ErrInternalServer = fmt.Errorf("internal server error")
		ErrBadRequest     = fmt.Errorf("bad request")
		ErrNotFound       = fmt.Errorf("not found")
	```