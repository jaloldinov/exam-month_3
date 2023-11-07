package v1

type SignInReq struct {
	Username string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignInResp struct {
	Token string `json:"token"`
}
