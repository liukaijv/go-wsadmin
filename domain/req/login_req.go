package req

import "errors"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginReq) Validate() error {

	if r.Username == "" {
		return errors.New("用户名不能为空")
	}

	if r.Password == "" {
		return errors.New("密码不能为空")
	}

	return nil
}
