package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Models"
	"go-chatbot/Pkg/Email"
	"go-chatbot/Pkg/MyRandom"
	"go.uber.org/zap"
)

// VerifiHandler 发送验证码
// @Summary 发送验证码接口
// @Description 发送验证码接口
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.VerifiEmail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /verificationCode [post]
func VerifiHandler(c *gin.Context) {
	// 数据绑定
	var verifEmail Models.VerifiEmail
	if err := c.ShouldBindJSON(&verifEmail); err != nil {
		zap.L().Error("Verifi Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 生成验证码并发送邮件
	verifiCode := MyRandom.RandomString(6)
	if err := Email.SendEmail(verifEmail.To, "验证码", verifEmail.Username, verifiCode); err != nil {
		zap.L().Error("Verifi Send Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeSendVerifiErr)
		return
	}
	// 将验证码存储到 Redis
	if err := Redis.RedisSet(verifEmail.To, verifiCode); err != nil {
		zap.L().Error("Redis Set Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}
	// 返回响应
	Models.ResponseSuccess(c, []string{})
}

// VerifiExam 验证验证码
// @Summary 验证验证码
// @Description 需要邮箱与验证码
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.VerifiExam false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /verificationCode [post]
func VerifiExam(c *gin.Context) {
	var verifExam Models.VerifiExam
	// 从 Redis 中检查验证码是否匹配
	if err := c.ShouldBindJSON(&verifExam); err != nil {
		zap.L().Error("Verifi Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	val, err := Redis.RedisFind(verifExam.Email)
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val == verifExam.VerifiCode {
		// 验证码匹配，返回成功响应
		Models.ResponseSuccess(c, []string{})
	} else {
		// 验证码不匹配，返回错误响应
		Models.ResponseError(c, Models.CodeVerifiErr)
	}
}
