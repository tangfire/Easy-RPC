package rpc

import (
	"encoding/gob"
	"reflect"
	"testing"
)

// 基础功能测试
func TestEncodeDecode(t *testing.T) {
	// 正常情况测试
	testCases := []struct {
		name    string
		input   RPCData
		wantErr bool
	}{
		{
			name: "basic_types",
			input: RPCData{
				Name: "Add",
				Args: []interface{}{1, 2},
			},
		},
		{
			name: "complex_args",
			input: RPCData{
				Name: "ProcessData",
				Args: []interface{}{"test", 3.14, []byte{0x01, 0x02}},
			},
		},
		{
			name:    "empty_data",
			input:   RPCData{},
			wantErr: false, // 空结构允许编解码
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 编码测试
			encoded, err := encode(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Encode() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}

			// 解码测试
			decoded, err := decode(encoded)
			if err != nil {
				t.Fatalf("Decode() error = %v", err)
			}

			// 深度比对
			if !reflect.DeepEqual(tc.input, decoded) {
				t.Errorf("Decoded data mismatch\nWant: %+v\nGot : %+v", tc.input, decoded)
			}
		})
	}
}

// 错误处理测试
func TestDecodeErrors(t *testing.T) {
	// 无效数据测试
	invalidData := []byte{0x01, 0x02, 0x03} // 伪造的无效gob数据

	_, err := decode(invalidData)
	if err == nil {
		t.Error("Expected error for invalid gob data, got nil")
	}

	// 空字节测试
	_, err = decode([]byte{})
	if err == nil {
		t.Error("Expected error for empty data, got nil")
	}
}

// 自定义类型测试（需要提前注册）
func TestCustomTypes(t *testing.T) {
	type CustomStruct struct {
		Field1 string
		Field2 int
	}
	gob.Register(CustomStruct{}) // 必须注册自定义类型[7](@ref)

	data := RPCData{
		Name: "CustomMethod",
		Args: []interface{}{CustomStruct{"test", 42}},
	}

	encoded, err := encode(data)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	decoded, err := decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	if _, ok := decoded.Args[0].(CustomStruct); !ok {
		t.Errorf("Type assertion failed, got %T", decoded.Args[0])
	}
}
