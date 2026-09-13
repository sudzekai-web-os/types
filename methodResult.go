package types

type HandlerResult struct {
	Data       any
	Error      error
	StatusCode int
}
