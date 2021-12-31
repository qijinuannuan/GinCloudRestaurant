package dao

import (
	"gincloudrestaurant/model"
	"gincloudrestaurant/tool"
	"log"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) QueryMemberById(id int64) *model.Member {
	var member model.Member
	if _, err := md.Where(" id = ? ", id).Get(&member); err != nil {
		return nil
	}
	return &member
}

func (md *MemberDao) UpdateMemberAvatar(userId int64, fileName string) int64 {
	member := model.Member{Avatar: fileName}
	result, err := md.Where(" id = ? ", userId).Update(&member)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return result
}

func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member
	password = tool.EncoderSha256(password)
	_, err := md.Where(" user_name = ? and password = ? ", name, password).Get(&member)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &member
}

func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode{
	var sms model.SmsCode
	if _, err := md.Where(" phone = ? and code = ? ", phone, code).Get(&sms); err != nil {
		log.Fatal(err.Error())
	}
	return &sms
}

func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where(" mobile = ? ", phone).Get(&member); err != nil {
		log.Fatal(err.Error())
	}
	return &member
}

func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return result
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
