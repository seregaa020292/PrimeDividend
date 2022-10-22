package models

const (
	ActiveStatus Status = "active"
	WaitStatus   Status = "wait"
	RemoveStatus Status = "remove"
)

type Status string

func (s Status) IsActive() bool {
	return s == ActiveStatus
}

func (s Status) IsWait() bool {
	return s == WaitStatus
}

func (s Status) IsRemove() bool {
	return s == RemoveStatus
}

func (s Status) String() string {
	return string(s)
}
