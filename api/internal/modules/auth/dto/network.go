package dto

import (
	"github.com/google/uuid"

	"primedivident/internal/models/app/public/model"
	"primedivident/internal/modules/auth/entity"
	"primedivident/internal/modules/auth/service/strategy/auth"
)

func ModelUserNetworksCreating(network entity.Network, userID uuid.UUID, strategy auth.Name) model.UserNetworks {
	return model.UserNetworks{
		UserID:     userID,
		ClientID:   network.ClientID,
		ClientType: strategy.String(),
	}
}

func EntityUserNetworksByModel(networks model.UserNetworks, email, name string) entity.Network {
	return entity.Network{
		ClientID:   networks.ClientID,
		ClientType: networks.ClientType,
		Email:      email,
		Name:       name,
	}
}
