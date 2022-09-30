package entity

type Network struct {
	ClientID   string
	ClientType string
	Email      string
	Name       string
}

func NewNetwork(id, email, name, strategy string) Network {
	return Network{
		ClientID:   id,
		ClientType: strategy,
		Email:      email,
		Name:       name,
	}
}

func (n Network) IsEmpty() bool {
	return n == (Network{})
}
