package authservice

import (
	"context"
)

func (s Service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	// TODO: check user status

	accessToken, err := s.tokenSvc.RefreshAccessToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
