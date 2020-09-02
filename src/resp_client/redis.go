package resp_client

import (
	"errors"
	"net/http"
	"encoding/json"
	my_redis "github.com/go-redis/redis"
	"context"
	// "fmt"
)

type RespRedis struct {
	Client *my_redis.Client
	ctx context.Context
	Ref string `json:"ref"`
}

func (resp *RespRedis) GetResp(r *http.Request) (string, error)  {
	err := json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return "", err
	}

	if err := resp.valid(); err != nil {
		return "", err
	}

	val, err := resp.Client.Get(resp.ctx, resp.Ref).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (resp *RespRedis) valid() error {
	if resp.Ref == "" {
		return errors.New("Param Ref is wrong")
	}
	return nil
}

func (resp *RespRedis) Init(param RespParam) error {
	resp.ctx = context.Background()
	resp.Client = my_redis.NewClient(&my_redis.Options{
        Addr:     param.Addr,
        Password: param.Password, 
        DB:       param.DB,
    })
	return nil
}