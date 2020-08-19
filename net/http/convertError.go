package http

import (
	"errors"
	netHttp "net/http"

	"github.com/plgd-dev/kit/coapconv"
	"github.com/plgd-dev/kit/grpcconv"
	coapStatus "github.com/plgd-dev/kit/net/coap/status"
	grpcStatus "google.golang.org/grpc/status"
)

type grpcErr interface {
	GRPCStatus() *grpcStatus.Status
}

// ErrToStatusWithDef converts err with default http.Status(for unknown conversion) to http.Status.
func ErrToStatusWithDef(err error, def int) int {
	if err == nil {
		return netHttp.StatusOK
	}
	var coapStatus coapStatus.Status
	if errors.As(err, &coapStatus) {
		return coapconv.ToHTTPCode(coapStatus.Message().Code, def)
	}
	var grpcErr grpcErr
	if errors.As(err, &grpcErr) {
		return grpcconv.ToHTTPCode(grpcErr.GRPCStatus().Code(), def)
	}
	return def
}

// ErrToStatus converts err to http.Status.
func ErrToStatus(err error) int {
	return ErrToStatusWithDef(err, netHttp.StatusInternalServerError)
}
