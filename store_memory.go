package ratelimiter

import (
	"context"
	"sync/atomic"
)

func (storeClient *MemoryStoreClient) Get(ctx context.Context, key string) (*int32, error) {
	_, exists := storeClient.buckets[key]
	if !exists {
		return nil, nil
	}
	value := storeClient.buckets[key].Load()
	return &value, nil
}

func (storeClient *MemoryStoreClient) Set(ctx context.Context, key string, value int32) error {
	_, exists := storeClient.buckets[key]
	if !exists {
		storeClient.buckets[key] = &atomic.Int32{}
	}

	storeClient.buckets[key].Store(value)
	return nil
}
