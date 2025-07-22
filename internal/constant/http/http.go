package http

// constant

// 20x
const (
	OK         = 200 // OK - 请求成功
	CREATED    = 201 // CREATED - 请求已成功并创建了新的资源
	ACCEPTED   = 202 // ACCEPTED - 请求已接受但尚未处理
	NO_CONTENT = 204 // NO_CONTENT - 请求成功，但没有返回内容
)

// 40x
const (
	BAD_REQUEST        = 400 // BAD_REQUEST - 请求无效
	UNAUTHORIZED       = 401 // UNAUTHORIZED - 未授权，身份验证失败
	FORBIDDEN          = 403 // FORBIDDEN - 禁止访问
	NOT_FOUND          = 404 // NOT_FOUND - 请求的资源未找到
	METHOD_NOT_ALLOWED = 405 // METHOD_NOT_ALLOWED - 请求方法不被允许
	CONFLICT           = 409 // CONFLICT - 资源冲突
)

// 50x
const (
	INTERNAL_SERVER_ERROR = 500 // INTERNAL_SERVER_ERROR - 服务器内部错误
	NOT_IMPLEMENTED       = 501 // NOT_IMPLEMENTED - 服务器不支持请求的功能
	BAD_GATEWAY           = 502 // BAD_GATEWAY - 网关错误
	SERVICE_UNAVAILABLE   = 503 // SERVICE_UNAVAILABLE - 服务不可用
)
