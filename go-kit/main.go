package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"learn-golang/go-kit/Endpoint"
	"learn-golang/go-kit/Services"
	"learn-golang/go-kit/Transport"
	"net/http"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Every(time.Second), 1)
	userService := Services.UserService{}
	userEndPoint := Endpoint.RateLimit(limiter)(Endpoint.GetUserEndPoint(userService))
	// 构造服务，实现http.handler并包装endpoint层
	serverHandler := httpTransport.NewServer(userEndPoint, Transport.DecodeUserRequest, Transport.EnCodeUserResponse)
	r := mux.NewRouter()

	//这种写法支持多种请求方法，访问Examp: http://localhost:8080/user/121便可以访问
	//r.Handle(`/user/{uid:\d+}`, serverHandler)

	//这种写法限定了请求只支持GET方法
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	// 监听端口，并且使用serverHandler处理随之而来的请求
	http.ListenAndServe(":8080", r)
}
