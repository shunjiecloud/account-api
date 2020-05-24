package schemas

type CreateUserRequest struct {
	// CaptchaId       string `json:"captcha_id" validate:"max=32"`
	// CaptchaSolution string `json:"captcha_solution" validate:"max=32"`
	Gender   int    `json:"gender" binding:"min=0,max=1"`
	Mail     string `json:"mail" binding:"email"`
	Phone    string `json:"phone" binding:"len=11"`
	Name     string `json:"name" binding:"min=1,max=16"`
	Password string `json:"password" binding:"len=32"`
}
