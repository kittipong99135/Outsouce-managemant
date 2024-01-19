package repository

import (
	"bytes"
	"net/http"
)

type AuthRepository interface {
	RepGetPwd(*bytes.Buffer) (*http.Request, error)
}
