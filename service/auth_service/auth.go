package service

import "osm/models"

type AuthService interface {
	SrvLogin(*models.Login_body) (*models.Login_response, error)
}
