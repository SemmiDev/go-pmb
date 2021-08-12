package interfaces

import (
	"context"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
)

type IRedisAuthRepository interface {
	CreateAuth(context.Context, string, *adapters.TokenDetails) error
	FetchAuth(context.Context, string) (string, error)
	DeleteRefresh(context.Context, string) error
	DeleteTokens(context.Context, *adapters.AccessDetails) error
}
