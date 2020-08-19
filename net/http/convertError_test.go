package http

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/plgd-dev/go-coap/v2/message"
	coapCodes "github.com/plgd-dev/go-coap/v2/message/codes"
	coapStatus "github.com/plgd-dev/kit/net/coap/status"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrToStatus(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ol", args: args{err: nil}, want: http.StatusOK},
		{name: "grpc", args: args{err: status.Error(codes.PermissionDenied, "grpc error")}, want: http.StatusForbidden},
		{
			name: "coap",
			args: args{
				err: coapStatus.Error(&message.Message{Code: coapCodes.Forbidden}, fmt.Errorf("coap error")),
			},
			want: http.StatusForbidden,
		},
		{name: "grpc", args: args{err: fmt.Errorf("unknown error")}, want: http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ErrToStatus(tt.args.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
