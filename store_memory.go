package ratelimiter

import (
	"context"
	"sync/atomic"
)

func (memoryStoreClient *MemoryStoreClient) Get(ctx context.Context, key string) (*int32, error) {
	_, exists := memoryStoreClient.buckets[key]
	if !exists {
		return nil, nil
	}
	value := memoryStoreClient.buckets[key].Load()
	return &value, nil
}

func (memoryStoreClient *MemoryStoreClient) Set(ctx context.Context, key string, value int32) error {
	_, exists := memoryStoreClient.buckets[key]
	if !exists {
		memoryStoreClient.buckets[key] = &atomic.Int32{}
	}

	memoryStoreClient.buckets[key].Store(value)
	return nil
}
