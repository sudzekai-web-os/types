package types

const (
	None  LogLevel = -1
	Debug LogLevel = iota
	Information
	Warning
	Error
	Critical
)

type LogLevel int

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Information:
		return "INFO"
	case Warning:
		return "WARN"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}
