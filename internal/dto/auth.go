package dto

type AuthSendOTPReq struct {
	Phone string `json:"phone"`
}

type AuthSendSellerOTPReq struct {
	Phone string `json:"phone"`
}

type AuthLoginRegisterReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRegisterRes struct {
	HasAccount bool     `json:"has_account"`
	Token      TokenRes `json:"token"`
}
