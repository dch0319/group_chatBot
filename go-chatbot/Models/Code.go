package Models

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParm
	CodeUserExist
	CodeEmailExist
	CodeUserNotExist
	CodeWrongPassword
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthHeaderIsNil
	CodeAuthHeaderValid
	CodeSendVerifiErr
	CodeVerifiNotFund
	CodeVerifiErr
	CodeIllegalityFriendReq
	CodeExceedMsg
	CodeWebSocketEnd
	CodeWebSocketReply
	CodeWebSocketReplyNoOne
	CodeWSSuccess
)

var CodeMesMap = map[ResCode]string{
	CodeSuccess:             "success",
	CodeInvalidParm:         "请求参数错误",
	CodeUserExist:           "用户已存在",
	CodeEmailExist:          "邮箱已存在",
	CodeUserNotExist:        "用户不存在",
	CodeWrongPassword:       "密码错误",
	CodeInvalidPassword:     "密码格式错误",
	CodeServerBusy:          "服务正忙",
	CodeAuthHeaderIsNil:     "需要登录后进行操作",
	CodeAuthHeaderValid:     "无效的Token",
	CodeSendVerifiErr:       "发送验证码失败",
	CodeVerifiNotFund:       "验证码过期或未发送",
	CodeVerifiErr:           "验证码错误",
	CodeIllegalityFriendReq: "非法发送好友请求",
	CodeExceedMsg:           "消息数量超过限制",
	CodeWebSocketEnd:        "连接中断",
	CodeWebSocketReply:      "对方在线应答",
	CodeWebSocketReplyNoOne: "对方不在线",
	CodeWSSuccess:           "websocket连接成功",
}

//此方法用于根据枚举值查询对应的提示文本
func (c ResCode) Msg() string {
	msg, ok := CodeMesMap[c]
	if !ok {
		msg = CodeMesMap[CodeServerBusy]
	}
	return msg
}
