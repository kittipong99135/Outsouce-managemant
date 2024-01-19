package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"osm/models"
	repository "osm/repository/auth_repo"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type authService struct {
	// authRepo repository.AuthRepository
	authRepo repository.AuthRepository
}

func NewAuthRepository(authRepo repository.AuthRepository) AuthService {
	return authService{authRepo: authRepo}
}

func (s authService) SrvLogin(login_body *models.Login_body) (*models.Login_response, error) {
	godotenv.Load(".env")

	validate := validator.New()
	err_validate := validate.Struct(login_body)
	if err_validate != nil {
		return nil, err_validate
	}

	login_request := models.Login_request{
		GrantType:    "password",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Username:     login_body.Username,
		Password:     login_body.Password,
	}

	json_conv, err := json.Marshal(login_request)
	if err != nil {
		return nil, errors.New("form login mid match")
	}

	buffer := bytes.NewBuffer(json_conv)

	request_call, err := s.authRepo.RepGetPwd(buffer)
	if err != nil {
		return nil, errors.New("Get pwd invalid")
	}

	request_call.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	request_client, err := client.Do(request_call)
	if err != nil {
		return nil, errors.New("not compile request client")
	}

	fmt.Println(request_client)

	defer request_client.Body.Close()

	result, err := ioutil.ReadAll(request_client.Body)

	var login_reqponse models.Login_response
	err = json.Unmarshal(result, &login_reqponse)
	if err != nil {
		return nil, errors.New("can't convert json format")
	}

	return &login_reqponse, nil
}
