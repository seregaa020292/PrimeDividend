package categorize

import (
	"log"

	"primedividend/api/internal/modules/auth/service/strategy/auth"
)

type Categorize struct {
	Passwords PasswordStrategies
	Networks  NetworkStrategies
}

func NewCategorize() Categorize {
	return Categorize{
		Passwords: make(PasswordStrategies),
		Networks:  make(NetworkStrategies),
	}
}

type maps[T any] map[auth.Name]T

func (m maps[T]) List() []T {
	networks := make([]T, 0, len(m))

	for _, strategy := range m {
		networks = append(networks, strategy)
	}

	return networks
}

func (m maps[T]) Get(name auth.Name) T {
	var (
		strategy T
		ok       bool
	)

	if strategy, ok = m[name]; ok {
		return strategy
	}

	return strategy
}

func (m maps[T]) Set(name auth.Name, strategy T) {
	if _, ok := m[name]; ok {
		log.Fatalln("strategy already exist")
	}

	m[name] = strategy
}
