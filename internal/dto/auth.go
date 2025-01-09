package dto

type AuthLoginRegisterReq struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

type LoginRegisterRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
