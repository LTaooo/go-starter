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
		return "操作成功"
	case BadRequest:
		return "请求参数错误"
	case Unauthorized:
		return "未授权"
	case Forbidden:
		return "无权限"
	case NotFound:
		return "资源未找到"
	case InternalError:
		return "服务器内部错误"
	default:
		return "未知错误"
	}
}
