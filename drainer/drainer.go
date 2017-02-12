package drainer

type Drainer interface {
	Drain(message string)
}
