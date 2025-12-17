package ui

type Focus int

const (
	FocusTree Focus = iota
	FocusConfig
	FocusDiff
)

func (f Focus) Next() Focus {
	return (f + 1) % 3
}
