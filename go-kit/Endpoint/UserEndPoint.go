package Endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"learn-golang/go-kit/Services"
	"log"
)

// UserRequest user请求格式
type UserRequest struct {
	Uid    int    `json:"uid"`
	Method string `json:"method"`
}

// UserResponse user响应格式
type UserResponse struct {
	Result string `json:"result"`
}

func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("429: too many request")
			}
			return next(ctx, request)
		}
	}
}

func GetUserEndPoint(userService Services.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := "nothings"
		if r.Method == "GET" {
			result = userService.GetName(uint(r.Uid))
		} else if r.Method == "DELETE" {
			err := userService.DelUser(uint(r.Uid))
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid为%d的用户已删除", r.Uid)
				log.Println(result)
			}
		}

		return UserResponse{Result: result}, nil
	}
}
