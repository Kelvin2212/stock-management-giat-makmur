package skeleton

import (
	"context"
	"errors"
	"skeleton/internal/entity"
	"skeleton/internal/entity/auth"
)

type Data interface {
}

type AuthSvc interface {
	CheckAuth(ctx context.Context, _token, code string) (auth.Auth, error)
}

type Service struct {
	data    Data
	authSvc AuthSvc
}

func New(data Data, authSvc AuthSvc) Service {
	return Service{
		data:    data,
		authSvc: authSvc,
	}
}

func (s Service) checkPermission(ctx context.Context, _permissions ...string) error {
	claims := ctx.Value(entity.ContextKey("claims"))
	if claims != nil {
		actions := claims.(entity.ContextValue).Get("permissions").(map[string]interface{})
		for _, action := range actions {
			permissions := action.([]interface{})
			for _, permission := range permissions {
				for _, _permission := range _permissions {
					if permission.(string) == _permission {
						return nil
					}
				}
			}
		}
	}
	return errors.New("401 unauthorized")
}
