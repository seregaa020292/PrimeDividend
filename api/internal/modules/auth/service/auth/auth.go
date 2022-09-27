package auth

type Auth interface {
	PasswordStrategy(key Key) PasswordStrategy
	NetworkStrategy(key Key) NetworkStrategy
	Verify(accessToken string) error
	Logout(refreshToken string) error
	Refresh(refreshToken string) (Tokens, error)
}

type auth struct {
	strategy   Strategy
	jwtTokens  JwtTokens
	repository TokenRepository
}

func NewAuth(
	strategy Strategy,
	jwtTokens JwtTokens,
	repository TokenRepository,
) Auth {
	return auth{
		strategy:   strategy,
		jwtTokens:  jwtTokens,
		repository: repository,
	}
}

func (a auth) PasswordStrategy(key Key) PasswordStrategy {
	return a.strategy.GetPassword(key)
}

func (a auth) NetworkStrategy(key Key) NetworkStrategy {
	return a.strategy.GetNetwork(key)
}

func (a auth) Verify(accessToken string) error {
	_, err := a.jwtTokens.ValidateAccessToken(accessToken)

	return err
}

func (a auth) Logout(refreshToken string) error {
	return nil
}

func (a auth) Refresh(refreshToken string) (Tokens, error) {
	return Tokens{}, nil
}
