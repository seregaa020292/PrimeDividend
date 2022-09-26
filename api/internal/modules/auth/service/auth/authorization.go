package auth

type Authorization interface {
	CheckAccess() error
	CheckRole() error
}
