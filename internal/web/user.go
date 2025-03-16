package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		passwordRegexPattern = `^(?=.*[a-zA-Z])(?=.*\d)[A-Za-z\d@$!%*?&]{8,}$`
	)

	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)

	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterUserServer(server *gin.Engine) {
	userServer := server.Group("/api/user")
	{
		userServer.POST("/signUp", u.SignUp)
		userServer.POST("/login", u.Login)
		userServer.POST("/profile", u.Profile)
		userServer.POST("/edit", u.Edit)

		// 这是测试接口忽略
		userServer.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
			})
		})
	}
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq

	// Bind 内部会调用 c.AbortWithError()，向客户端返回错误响应，并停止后续中间件或处理器的执行。
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// password 和 confirmPassword 是否一致
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusOK, "两次密码输入不一致")
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
	}
	if !ok {
		ctx.JSON(http.StatusOK, "邮箱格式无效")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
	}
	if !ok {
		ctx.JSON(http.StatusOK, "密码格式无效,密码总长度至少为8个字符，且只能包含字母、数字和@$!%*?&等特殊字符")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"token": "token",
		},
	})
}

func (*UserHandler) Login(ctx *gin.Context) {

}

func (*UserHandler) Profile(ctx *gin.Context) {

}

func (*UserHandler) Edit(ctx *gin.Context) {

}
