package service

import (
	"encoding/json"
	"fmt"
	"gincloudrestaurant/dao"
	"gincloudrestaurant/model"
	"gincloudrestaurant/param"
	"gincloudrestaurant/tool"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {

}

func (ms *MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	md := dao.MemberDao{Orm: tool.DbEngine}
	return md.QueryMemberById(int64(id))
}

func (ms *MemberService) UploadAvatar(userId int64, fileName string) string {
	memberDao := dao.MemberDao{Orm: tool.DbEngine}
	result := memberDao.UpdateMemberAvatar(userId, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}

// Login 用户登录
func (ms *MemberService) Login(name string, password string) *model.Member {
	md := dao.MemberDao{Orm: tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}

	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user
}

// SmsLogin 用户手机号+验证码登录
func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member{
	md := dao.MemberDao{Orm: tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}

	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}

	user := model.Member{}
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user
}

func (ms *MemberService) SendCode(phone string)  bool {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	log.Println(response)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if response.Code == "OK" {
		smsCode := model.SmsCode{
			Phone: phone,
			Code: code,
			BizId: response.BizId,
			CreateTime: time.Now().Unix(),
		}
		memberDao := dao.MemberDao{Orm: tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	return false
}

type ResponseSms struct {
	BizId string `json:"biz_id"`
	Code string `json:"code"`
}

func (ms *MemberService) SendCodeTest(phone string)  bool {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	response := ResponseSms {
		BizId: tool.RandomRequestID(),
		Code: "OK",
	}
	if response.Code == "OK" {
		smsCode := model.SmsCode{
			Phone: phone,
			Code: code,
			BizId: response.BizId,
			CreateTime: time.Now().Unix(),
		}
		memberDao := dao.MemberDao{Orm: tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	return false
}