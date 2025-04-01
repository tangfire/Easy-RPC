package rpc

import (
	"fmt"
	"net"
	"reflect"
)

// 声明服务端
type Server struct {
	// 地址
	addr string
	// map 用于维护关系的
	funcs map[string]reflect.Value
}

// 构造方法
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// 服务端需要一个注册Register
// 第一个参数函数名，第二个传入真正的函数
func (s *Server) Register(rpcName string, f interface{}) {
	// 维护一个map
	// 若map已经有键了
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// 若map中没值，则将映射加入map，用于调用
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

// 服务端等待调用的方法
func (s *Server) Run() {
	// 监听
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s err :%v", s.addr, err)
		return
	}
	for {
		// 服务端循环等待调用
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := NewSession(conn)
		// 使用RPC方式读取数据
		b, err := serSession.Read()
		if err != nil {
			return
		}
		// 数据解码
		rpcData, err := decode(b)
		if err != nil {
			return
		}
		// 根据读到的name，得到要调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在", rpcData.Name)
			return
		}
		// 遍历解析客户端传来的参数,放切片里
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法
		// 返回Value类型，用于给客户端传递返回结果,out是所有的返回结果
		out := f.Call(inArgs)
		// 遍历out ，用于返回给客户端，存到一个切片里
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		// 数据编码，返回给客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respRPCData)
		if err != nil {
			return
		}
		// 将服务端编码后的数据，写出到客户端
		err = serSession.Write(bytes)
		if err != nil {
			return
		}
	}
}
