package cli

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/spf13/pflag"
// )

// func TestValue_Set(t *testing.T) {
// 	type want struct {
// 		value any
// 		err   error
// 	}
// 	tests := []struct {
// 		name  string
// 		value pflag.Value
// 		arg   string
// 		want  want
// 	}{
// 		{
// 			name:  "string",
// 			value: &Value[string]{},
// 			arg:   "test",
// 			want: want{
// 				value: "test",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "bool true",
// 			value: &Value[bool]{},
// 			arg:   "true",
// 			want: want{
// 				value: true,
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "bool false",
// 			value: &Value[bool]{},
// 			arg:   "false",
// 			want: want{
// 				value: false,
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "bool none",
// 			value: &Value[bool]{},
// 			arg:   "",
// 			want: want{
// 				value: true,
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "int32",
// 			value: &Value[int32]{},
// 			arg:   "123",
// 			want: want{
// 				value: "123",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "int64",
// 			value: &Value[int64]{},
// 			arg:   "123",
// 			want: want{
// 				value: "123",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "uint32",
// 			value: &Value[uint32]{},
// 			arg:   "123",
// 			want: want{
// 				value: "123",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "uint64",
// 			value: &Value[uint64]{},
// 			arg:   "123",
// 			want: want{
// 				value: "123",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "float32",
// 			value: &Value[float32]{},
// 			arg:   "1.1",
// 			want: want{
// 				value: "1.1",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "float64",
// 			value: &Value[float64]{},
// 			arg:   "1.1",
// 			want: want{
// 				value: "1.1",
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "[]byte base64",
// 			value: &Value[[]byte]{},
// 			arg:   "dGVzdA==",
// 			want: want{
// 				value: []byte("test"),
// 				err:   nil,
// 			},
// 		},
// 		{
// 			name:  "struct{}",
// 			value: &Value[struct{}]{},
// 			arg:   "",
// 			want: want{
// 				value: struct{}{},
// 				err:   ErrUnimplemented,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.value.Set(tt.arg); err != tt.want.err {
// 				t.Errorf("Value.Set() error = %v", err)
// 			}
// 			if fmt.Sprintf("%v", tt.want.value) != tt.value.String() {
// 				t.Errorf("Value.Set() expected type %T value %v, got value %s", tt.want.value, tt.want.value, tt.value.String())
// 			}
// 		})
// 	}
// }
