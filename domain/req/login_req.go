package req

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
