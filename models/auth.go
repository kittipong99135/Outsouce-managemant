package models

// Redirect struct
type Redirect_response struct {
	Key_1 string `json:"Key_1"`
	Key_2 string `json:"Key_2"`
}

// Login struct
type Login_body struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login_request struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type Login_response struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountID      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}
