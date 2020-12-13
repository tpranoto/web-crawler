package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestResponse_WriteResponse(t *testing.T) {
	writer := httptest.NewRecorder()

	type fields struct {
		Header    HeaderResponse
		Data      interface{}
		startTime time.Time
	}
	type args struct {
		w   http.ResponseWriter
		msg interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "valid",
			fields: fields{},
			args: args{
				w:   writer,
				msg: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &Response{
				Header:    tt.fields.Header,
				Data:      tt.fields.Data,
				startTime: tt.fields.startTime,
			}
			res.WriteResponse(tt.args.w, tt.args.msg)
		})
	}
}

func TestResponse_WriteError(t *testing.T) {
	writer := httptest.NewRecorder()

	type fields struct {
		Header    HeaderResponse
		Data      interface{}
		startTime time.Time
	}
	type args struct {
		w    http.ResponseWriter
		code int
		msg  interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "valid",
			args: args{
				w:    writer,
				code: http.StatusOK,
				msg:  "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &Response{
				Header:    tt.fields.Header,
				Data:      tt.fields.Data,
				startTime: tt.fields.startTime,
			}
			res.WriteError(tt.args.w, tt.args.code, tt.args.msg)
		})
	}
}
