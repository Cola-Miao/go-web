package model

type SignupForm struct {
	Name  string `form:"name" binding:"min=1,max=16"`
	Pwd   string `form:"pwd" binding:"min=8,max=32"`
	RePwd string `form:"rePwd" binding:"eqfield=Pwd"`
}

type LoginForm struct {
	Name string `form:"name" binding:"required,max=16"`
	Pwd  string `form:"pwd" binding:"min=8,max=32"`
}
