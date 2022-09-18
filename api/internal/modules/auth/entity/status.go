package entity

const (
	Wait   Status = "wait"
	Active Status = "active"
)

type Status string

func (s Status) IsWait() bool {
	return s == Wait
}

func (s Status) IsActive() bool {
	return s == Active
}

func (s Status) String() string {
	return string(s)
}
