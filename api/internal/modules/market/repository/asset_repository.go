package repository

import (
	"context"
	"fmt"

	"primedivident/pkg/db/redis"
)

type AssetRepository interface {
	GetAsset(key string) error
	SaveAsset(key string, value any) error
	RemoveAsset(keys ...string) error
}

type assetRepository struct {
	ctx context.Context
	db  *redis.Redis
}

func NewAssetRepository(db *redis.Redis) AssetRepository {
	return assetRepository{
		ctx: context.Background(),
		db:  db,
	}
}

func (r assetRepository) GetAsset(key string) error {
	_, err := r.db.Get(r.ctx, r.assetPrefix(key)).Bytes()
	if err != nil {
		return err
	}

	return nil
}

func (r assetRepository) SaveAsset(key string, value any) error {
	return r.db.Set(r.ctx, r.assetPrefix(key), value, 0).Err()
}

func (r assetRepository) RemoveAsset(keys ...string) error {
	for i, key := range keys {
		keys[i] = r.assetPrefix(key)
	}
	return r.db.Del(r.ctx, keys...).Err()
}

func (r assetRepository) assetPrefix(key string) string {
	return fmt.Sprintf("asset:%s", key)
}
