package resp

type LoginResp struct {
	Token  string `json:"token"`
	Expire int    `json:"expire"`
}
