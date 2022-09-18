package entity

const (
	Wait   Status = "wait"
	Active Status = "active"
)

type Status string

func (s Status) isWait() bool {
	return s == Wait
}

func (s Status) isActive() bool {
	return s == Active
}

func (s Status) String() string {
	return string(s)
}
