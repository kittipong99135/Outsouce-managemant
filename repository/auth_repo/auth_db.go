package repository

import (
	"bytes"
	"net/http"
)

type authRepository struct {
}

func NewAuthRepository() AuthRepository {
	return authRepository{}
}
func (r authRepository) RepGetPwd(buffer *bytes.Buffer) (*http.Request, error) {
	request, err := http.NewRequest("POST", "https://one.th/api/oauth/getpwd", buffer)
	if err != nil {
		return nil, err
	}

	return request, nil
}
