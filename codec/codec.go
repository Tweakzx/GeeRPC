package codec

import "io"

type Header struct {
	ServiceMethod string // 服务名和方法名
	Seq           uint64 // 请求的序号
	Error         string // 错误信息
}

//对消息进行编解码的接口
type Codec interface {
	io.Closer //需要实现一个close（）函数
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
	NewCodecFuncMap[JsonType] = NewJsonCodec
}
