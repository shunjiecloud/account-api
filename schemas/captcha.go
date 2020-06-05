package schemas

type CreateUserRequest struct {
	CaptchaId       string `json:"captcha_id" binding:"max=32"`
	CaptchaSolution string `json:"captcha_solution" binding:"max=32"`
	Name            string `json:"name" binding:"min=3,max=16"`
	Password        string `json:"password"`
	Mail            string `json:"mail" binding:"email"`
	Phone           string `json:"phone" binding:"len=11"`
	Gender          int    `json:"gender" binding:"min=0,max=1"`
	Avatar          string `json:"avatar" binding:"max=256"`
}
