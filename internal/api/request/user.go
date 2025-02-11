package request

import (
	"errors"

	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

var (
	ErrBadRequest = errors.New("empty name")
)

func ID(req *userv1.GetByIdRequest) int {
	return int(req.GetId())
}

func Name(req *userv1.GetByNameRequest) (string, error) {
	name := req.GetName()
	if name == "" {
		return "", ErrBadRequest
	}

	return name, nil
}
