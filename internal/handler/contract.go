package handler

import "context"

type UserService interface {
	RegisterUser(ctx context.Context, req *Request[CreateUserReq]) (*Response[*CreateUserRes], error)
	Login(ctx context.Context, req *Request[LoginReq]) (*Response[*LoginRes], error)
	GetUserById(ctx context.Context, req *Request[GetUserByIdReq]) (*Response[*GetUserByIdRes], error)
}

type Validator interface {
	Validate() error
}
