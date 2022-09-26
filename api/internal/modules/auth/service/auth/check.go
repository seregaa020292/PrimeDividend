package auth

type Check interface {
	Access() error
	Role() error
}
