package controller

import (
	"encoding/json"
	"gincloudrestaurant/model"
	"gincloudrestaurant/param"
	"gincloudrestaurant/service"
	"gincloudrestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.SendSmsCode)
	engine.POST("/api/login_sms", mc.SmsLogin)
	engine.GET("api/captcha", mc.Captcha)
	engine.POST("api/verifycha", mc.VerifyCaptcha)
	engine.POST("/api/login_pwd", mc.NameLogin)
	engine.POST("api/upload/avatar", mc.UploadAvatar)
	engine.POST("/api/userinfo", mc.UserInfo)
}

func (mc *MemberController) UserInfo(ctx *gin.Context) {
	cookie, err := tool.CookieAuth(ctx)
	if err != nil {
		ctx.Abort()
		tool.Failed(ctx, "login first")
		return
	}
	memberService := service.MemberService{}
	member := memberService.GetUserInfo(cookie.Value)
	if member != nil {
		tool.Success(ctx, map[string]interface{}{
			"id": member.Id,
			"user_name": member.UserName,
			"mobile": member.Mobile,
			"register_time": member.RegisterTime,
			"avatar": member.Avatar,
			"balance": member.Balance,
			"city": member.City,
		})
		return
	}
	tool.Failed(ctx, "get user info failed")
}

// UploadAvatar 头像上传
func (mc *MemberController) UploadAvatar(ctx *gin.Context) {
	userId := ctx.PostForm("user_id")
	log.Println(userId)
	file, err := ctx.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Failed(ctx, "param parse failed")
		return
	}

	session := tool.GetSession(ctx, "user_" + userId)
	if session == nil {
		tool.Failed(ctx, "param error")
		return
	}
	var member model.Member
	json.Unmarshal(session.([]byte), &member)

	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(ctx, "avatar update failed")
		return
	}

	memberService := service.MemberService{}
	path := memberService.UploadAvatar(member.Id, fileName[1:])
	if path != "" {
		tool.Success(ctx, "http://localhost:8080" + path)
		return
	}
	tool.Failed(ctx, "upload failed")
}


// SendSmsCode http://localhost:8080/api/sencode?phone=13122225079
func (mc *MemberController) SendSmsCode(ctx *gin.Context) {
	phone, exit := ctx.GetQuery("phone")
	if !exit {
		tool.Failed(ctx, "param parse failed")
		return
	}
	ms := service.MemberService{}
	// SendCode
	isSend := ms.SendCodeTest(phone)
	if isSend {
		tool.Success(ctx, "send success")
		return
	}
	tool.Failed(ctx, "send failed")
}

func (mc *MemberController) SmsLogin(ctx *gin.Context) {
	var smsLoginParam param.SmsLoginParam
	if err := tool.Decode(ctx.Request.Body, &smsLoginParam); err != nil {
		tool.Failed(ctx, "param parse failed")
		return
	}
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		sess, _ := json.Marshal(member)
		err := tool.SetSession(ctx, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(ctx, "save session error, login failed")
		}
		ctx.SetCookie("cookie_user", strconv.Itoa(int(member.Id)), 10*60, "/", "localhost", true, true)
		tool.Success(ctx, member)
		return
	}
	tool.Failed(ctx, "login failed")
	return
}

// Captcha 生成验证码
func (mc *MemberController) Captcha(ctx *gin.Context) {
	//todo 生成验证码，并返回客户端
	tool.GenerateCaptcha(ctx)
}

// VerifyCaptcha 验证验证码正确
func (mc *MemberController) VerifyCaptcha(ctx *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(ctx.Request.Body, &captcha)
	if err != nil {
		tool.Failed(ctx, "param parse failed")
		return
	}
	if result := tool.VerifyCaptcha(captcha.Id, captcha.VerifyValue); result {
		log.Println("verify success")
		tool.Success(ctx, "verify success")
	} else {
		log.Println("verify failed")
		tool.Failed(ctx, "verify failed")
	}
}

// NameLogin 用户名+密码、验证码登录
func (mc *MemberController) NameLogin(ctx *gin.Context) {
	var loginParam param.LoginParam
	err := tool.Decode(ctx.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(ctx, "param parse failed")
		return
	}
	validate := tool.VerifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(ctx, "验证码不正确，重新验证")
		return
	}
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		sess, _ := json.Marshal(member)
		err := tool.SetSession(ctx, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(ctx, "save session error, login failed")
		}
		tool.Success(ctx, &member)
		return
	}
	tool.Failed(ctx, "login failed")
}