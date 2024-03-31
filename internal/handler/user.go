package handler

type UserHandler struct {
	UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService}
}
