package auth

type Authorization interface {
	checkAccess() error
	checkRole() error
}
