package auth

import (
	"context"
	"net/http"
	"skeleton/internal/entity/auth"
	"skeleton/pkg/errors"
)

// CheckAuth ...
func (d Data) CheckAuth(ctx context.Context, _token, code string) (auth.Auth, error) {
	authResponse := auth.Auth{}
	endpoint := d.baseURL + "/checkrights"
	body := map[string]interface{}{
		"code": code,
	}

	headers := make(http.Header)
	headers.Set("Authorization", _token)
	headers.Set("Content-Type", "application/json")

	_, err := d.client.PostJSON(ctx, endpoint, headers, body, &authResponse)
	if err != nil {
		return authResponse, errors.Wrap(err, "[DATA][CheckAuth]")
	}

	return authResponse, nil
}
