package handler

import (
	"errors"

	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

var errIsNotLiveStreaming = errors.New("this video is not live streaming")

func withErrInfo(code connect.Code, err error, msg string) *connect.Error {
	d, de := connect.NewErrorDetail(&errdetails.ErrorInfo{
		Reason: msg,
	})
	if de != nil {
		panic(de)
	}

	ce := connect.NewError(code, err)
	ce.AddDetail(d)

	return ce
}
