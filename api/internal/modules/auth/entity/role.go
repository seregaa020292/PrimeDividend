package entity

const (
	AdminRole   Role = "admin"
	ManagerRole Role = "manager"
	UserRole    Role = "user"
	GuestRole   Role = "guest"
)

type Role string

func (r Role) String() string {
	return string(r)
}
