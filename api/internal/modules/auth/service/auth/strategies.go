package auth

import "log"

type (
	MapStrategy[T any] map[Key]T
	PasswordStrategies = MapStrategy[PasswordStrategy]
	NetworkStrategies  = MapStrategy[NetworkStrategy]
)

func (m MapStrategy[T]) list() []T {
	networks := make([]T, 0, len(m))

	for _, strategy := range m {
		networks = append(networks, strategy)
	}

	return networks
}

func (m MapStrategy[T]) get(name Key) T {
	var (
		strategy T
		ok       bool
	)

	if strategy, ok = m[name]; ok {
		return strategy
	}

	return strategy
}

func (m MapStrategy[T]) set(name Key, strategy T) {
	if _, ok := m[name]; ok {
		log.Fatalln("strategy already exist")
	}

	m[name] = strategy
}
