package role

type Check interface {
	Access() error
	IsAdmin() error
	IsUser() error
}
