package entity

const (
	WaitStatus   Status = "wait"
	ActiveStatus Status = "active"
)

type Status string

func (s Status) IsWait() bool {
	return s == WaitStatus
}

func (s Status) IsActive() bool {
	return s == ActiveStatus
}

func (s Status) String() string {
	return string(s)
}
