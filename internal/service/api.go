package service

import (
	"context"
	"fmt"

	"cachedapi/pkg/cache"
)

var ErrCacheMiss = fmt.Errorf("cache miss: requested key not found")

type ApiService struct {
	cache *cache.Client
}

func NewApiService(cache *cache.Client) *ApiService {
	return &ApiService{
		cache: cache,
	}
}

func (s *ApiService) GetCache(ctx context.Context, key string) ([]byte, error) {
	exists, err := s.cache.Exists(ctx, key)
	if err != nil || !exists {
		return nil, ErrCacheMiss
	}

	data, err := s.cache.Get(ctx, key)
	if err != nil {
		return nil, ErrCacheMiss
	}

	return data, nil
}

func (s *ApiService) SetCache(ctx context.Context, key string, data []byte) error {
	return s.cache.Set(ctx, key, data)
}
