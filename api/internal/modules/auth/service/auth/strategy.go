package auth

type Key string

const (
	Email  Key = "email"
	Vk     Key = "vk"
	Ok     Key = "ok"
	Yandex Key = "yandex"
)

type (
	Strategy interface {
		GetPassword(key Key) PasswordStrategy
		GetNetwork(key Key) NetworkStrategy
		SetPassword(key Key, strategy PasswordStrategy)
		SetNetwork(key Key, strategy NetworkStrategy)
	}
	NetworkStrategy interface {
		Callback(state string) string
		Login(code string) (Tokens, error)
	}
	PasswordStrategy interface {
		Login(identify, password string) (Tokens, error)
	}
)

type strategy struct {
	passwords PasswordStrategies
	networks  NetworkStrategies
}

func NewStrategy() Strategy {
	return &strategy{
		passwords: make(PasswordStrategies),
		networks:  make(NetworkStrategies),
	}
}

func (s *strategy) GetPassword(key Key) PasswordStrategy {
	return s.passwords.get(key)
}

func (s *strategy) GetNetwork(key Key) NetworkStrategy {
	return s.networks.get(key)
}

func (s *strategy) SetPassword(key Key, strategy PasswordStrategy) {
	s.passwords.set(key, strategy)
}

func (s *strategy) SetNetwork(key Key, strategy NetworkStrategy) {
	s.networks.set(key, strategy)
}
