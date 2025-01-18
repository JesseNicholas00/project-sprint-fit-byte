package user

import "context"

type UserService interface {
	RegisterUser(ctx context.Context, req AuthenticationUserReq, res *AuthenticationUserRes) error
	LoginUser(ctx context.Context, req AuthenticationUserReq, res *AuthenticationUserRes) error
	GetSessionFromToken(ctx context.Context, req GetSessionFromTokenReq, res *GetSessionFromTokenRes) error
	FindUser(ctx context.Context, userID string, res *FindUserRes) error
	UpdateUser(ctx context.Context, userID string, req UpdateUserReq, res *UpdateUserRes) error
}
