package entity

type Network struct {
	Identity string
	Strategy string
	Email    string
	Name     string
}

func NewNetwork(id, email, name, strategy string) Network {
	return Network{
		Identity: id,
		Strategy: strategy,
		Email:    email,
		Name:     name,
	}
}

func (n Network) IsEmpty() bool {
	return n == (Network{})
}
