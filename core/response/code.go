package response

type Code int

const (
	// 2. 成功响应码
	OK Code = 200

	// 3. 客户端请求错误
	BadRequest   Code = 400
	Unauthorized Code = 401
	Forbidden    Code = 403
	NotFound     Code = 404

	// 4. 服务端错误
	InternalError Code = 500
)

// 6. 获取响应码对应的默认消息
func (c Code) Message() string {
	switch c {
	case OK:
		return "success"
	case BadRequest:
		return "bad request"
	case Unauthorized:
		return "unauthorized"
	case Forbidden:
		return "forbidden"
	case NotFound:
		return "not found"
	case InternalError:
		return "internal error"
	default:
		return "unknown error"
	}
}
