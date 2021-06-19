package entity

import "context"

// Counter represents counter for
type Counter struct {
	StudentID   uint32 `bson:"id" json:"id"`
}

// CounterRepository repo for counter
type CounterRepository interface {
	Get(ctx context.Context, collectionName string, identifier string) (uint32, error)
}
