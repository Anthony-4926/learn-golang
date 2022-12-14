package Transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"learn-golang/go-kit/Endpoint"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if _uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(_uid)
		return Endpoint.UserRequest{Uid: uid, Method: r.Method}, nil
	}
	return nil, errors.New("参数错误")
}

func EnCodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
