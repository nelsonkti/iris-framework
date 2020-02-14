package Exceptions

type ErrorCode struct {
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
	Error string `json:"error"`
}

// ErrorInsertDatabase 1-插入数据库失败
func ErrorInsertDatabase(err error) ErrorCode {
	return buildError(1, "插入数据库出错", err.Error())
}

// ErrorQueryDatabase 2-查询数据库失败
func ErrorQueryDatabase(err error) ErrorCode {
	return buildError(2, "查询数据库失败", err.Error())
}

// ErrorUpdateDatabase 3-更新数据库失败
func ErrorUpdateDatabase(err error) ErrorCode {
	return buildError(3, "更新数据库失败", err.Error())
}

// ErrorDeleteDatabase 4-更新数据库失败
func ErrorDeleteDatabase(err error) ErrorCode {
	return buildError(4, "删除数据库失败", err.Error())
}

// ErrorIncompleteData 5-数据不完整
func ErrorIncompleteData(err error) ErrorCode {
	return buildError(5, "数据不完整", err.Error())
}

// ErrorVerification 6-数据检验失败
func ErrorVerification(err error) ErrorCode {
	return buildError(6, "数据检验失败", err.Error())
}

// ErrorBuildJWT 7-生成凭证错误
func ErrorBuildJWT(err error) ErrorCode {
	return buildError(7, "生成凭证错误", err.Error())
}

// ErrorUnauthorized 8-未认证登录
func ErrorUnauthorized(err error) ErrorCode {
	return buildError(8, "未认证登录", err.Error())
}

func buildError(code int64, msg string, detail string) ErrorCode {
	return ErrorCode{
		Code:   code,
		Msg:    msg,
		Error: detail,
	}
}