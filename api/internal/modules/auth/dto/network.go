package dto

import (
	"github.com/google/uuid"

	"primedividend/api/internal/models/app/public/model"
	"primedividend/api/internal/modules/auth/entity"
	"primedividend/api/internal/modules/auth/service/strategy/auth"
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
