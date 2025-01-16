package user

import "context"

type UserService interface {
	RegisterUser(ctx context.Context, req AuthenticationUserReq, res *AuthenticationUserRes) error
}
