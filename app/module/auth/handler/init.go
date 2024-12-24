package handler

import (
	"github.com/Brotiger/per-painted_poker-backend/app/module/auth/service"
	sharedService "github.com/Brotiger/per-painted_poker-backend/app/shared/service"
)

type Auth struct {
	AuthService         *service.Auth
	RefreshTokenService *service.RefreshToken
	SharedTokenService  *sharedService.Token
}

func NewAuth() *Auth {
	return &Auth{
		AuthService:         service.NewAuth(),
		RefreshTokenService: service.NewRefreshToken(),
		SharedTokenService:  sharedService.NewToken(),
	}
}