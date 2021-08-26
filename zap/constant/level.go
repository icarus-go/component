package constant

type Level string

const (
	Debug  Level = "debug"
	Info   Level = "info"
	Warn   Level = "warn"
	Error  Level = "error"
	Dpanic Level = "dpanic"
	Panic  Level = "panic"
	Fatal  Level = "fatal"
)

func (l Level) Value() string {
	return string(l)
}
