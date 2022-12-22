package Transport

import (
	"context"
	"encoding/json"
	"errors"
	"learn-golang/go-kit/Endpoint"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	_uid := r.URL.Query().Get("uid")
	if uid, ok := strconv.Atoi(_uid); ok == nil {
		return Endpoint.UserRequest{Uid: uid, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

func EnCodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
