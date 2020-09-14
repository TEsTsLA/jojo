package util

import "net/http"

type BaseError struct {
	error
	Code ErrorCode `json:"code"`
	Msg  string    `json:"msg"`
}
// 错误码
type ErrorCode int

func (e *BaseError) Error() string {
	return e.Msg
}



const (
	SUCCESS     ErrorCode = http.StatusOK
	FALL        ErrorCode = http.StatusBadRequest
	NOT_LOGIN   ErrorCode = http.StatusUnauthorized
	VERIFY_FALL ErrorCode = 412
)

// 表单错误
var (
	FormParamErr = BaseError{Code: FALL, Msg: "请求参数有误"}
)

// 登录和身份验证相关错误
var (
	NotLoginError       = BaseError{Code: NOT_LOGIN, Msg: "没有登录"}
	UserNotExitError    = BaseError{Code: FALL, Msg: "用户不存在"}
	UserPasswordError   = BaseError{Code: FALL, Msg: "账号或密码错误"}
	UserRoleVerifyError = BaseError{Code: VERIFY_FALL, Msg: "用户身份验证失败"}
)

// 地址相关
var (
	AddressNotFound = BaseError{Code: FALL, Msg: "没有找到对应地址"}
	AddressInfoErr  = BaseError{Code: FALL, Msg: "地址信息有误"}
)

// 产品相关错误
var (
	ProductNotFound    = BaseError{Code: FALL, Msg: "没有找到相关产品"}
	ProductPkgNotFound = BaseError{Code: FALL, Msg: "没有找到相关套餐"}
	OutOfStock         = BaseError{Code: FALL, Msg: "库存不足"}
)

// 订单相关错误
var (
	OrderCreateErr = BaseError{Code: FALL, Msg: "创建订单时发生错误"}
)
