package Controler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

func ChangeUserImg(c *gin.Context)  {
	var p Models.ParmChangeImg
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	err := Logic.ChangeUserImg(&p)
	//业务处理
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

func ChangeUser(c *gin.Context)  {
	var p Models.ParmChange
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	err := Logic.ChangeUser(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

func ChangeNick(c *gin.Context)  {
	var p Models.ParmFriendRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	fmt.Println(p)
	q := Models.ParmSerchIsFriend{
		UserId:   p.UserId,
		FriendId: p.FriendId,
	}
	friendRes, err := Logic.SerchIsFriend(&q)
	if err != nil {
		if Models.ErrorIs(err, Models.ErrorWrongPassword) || Models.ErrorIs(err, Models.ErrorUserNotExit) {
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err.Error())
			return
		}
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	if friendRes.State != 0{
		zap.L().Error("SomeOne Illegality Send Friend Request:", zap.Error(err))
		Models.ResponseError(c, Models.CodeIllegalityFriendReq)
		return
	}
	//业务处理
	//state 0为好友 1为申请中 2为非好友

	if err := Logic.ChangeNick(&p); err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

func ChangeEmail(c *gin.Context)  {
	var p Models.ParmChangeEmail
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	fmt.Println(p)
	q := Models.VerifiExam{
		VerifiCode: p.VerifiCode,
		Email: p.NewEmail,
	}
	val, err := Redis.RedisFind(q.Email)
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val != q.VerifiCode {
		Models.ResponseError(c, Models.CodeVerifiErr)
		return
	}
	//业务处理
	s := Models.ParmChange{
		UserId: p.UserId,
		Data: p.NewEmail,
		Type: "email",
		Psw: "",
	}
	err = Logic.ChangeUser(&s)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

func ChangePsw(c *gin.Context)  {
	var p Models.ParmChangePsw
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	q := Models.VerifiExam{
		VerifiCode: p.VerifiCode,
		Email: p.Email,
	}
	val, err := Redis.RedisFind(q.Email)
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val != q.VerifiCode {
		Models.ResponseError(c, Models.CodeVerifiErr)
		return
	}
	if len(p.Password) < 8 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码过短")
		return
	}
	if len(p.RePassword) < 8 || p.RePassword != p.Password {
		Models.ResponseErrorWithMsg(c, Models.CodeWrongPassword, "请重新确认密码")
		return
	}
	//业务处理
	s := Models.ParmChange{
		UserId: p.UserId,
		Data: p.Password,
		Type: "password",
		Psw: "",
	}
	err = Logic.ChangeUser(&s)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}
