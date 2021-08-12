package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/database"
	"time"
)

type RedisAuthRepository struct {
	database.RedisConnection
}

func NewRedisAuthRepository(redisConn *database.RedisConnection) interfaces.IRedisAuthRepository {
	return &RedisAuthRepository{*redisConn}
}

func (r *RedisAuthRepository) CreateAuth(c context.Context, userid string, td *adapters.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := r.Client.Set(c, td.TokenUuid, userid, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := r.Client.Set(c, td.RefreshUuid, userid, rt.Sub(now)).Result()
	if err != nil {
		return err
	}

	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

func (r *RedisAuthRepository) FetchAuth(c context.Context, tokenUUID string) (string, error) {
	userid, err := r.Client.Get(c, tokenUUID).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (r *RedisAuthRepository) DeleteRefresh(c context.Context, refreshUuid string) error {
	//delete refresh token
	deleted, err := r.Client.Del(c, refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}

func (r *RedisAuthRepository) DeleteTokens(c context.Context, details *adapters.AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", details.TokenUuid, details.UserId)
	//delete access token
	deletedAt, err := r.Client.Del(c, details.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := r.Client.Del(c, refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}
