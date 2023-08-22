package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
)

func (h *GeneralAPIHandler) ConnectToMeetingStream(
	ctx context.Context,
	req *connect.Request[emoine_rv1.ConnectToMeetingStreamRequest],
	stream *connect.ServerStream[emoine_rv1.ConnectToMeetingStreamResponse],
) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendComment(ctx context.Context, req *connect.Request[emoine_rv1.SendCommentRequest]) (*connect.Response[emoine_rv1.SendCommentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendReaction(ctx context.Context, req *connect.Request[emoine_rv1.SendReactionRequest]) (*connect.Response[emoine_rv1.SendReactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}
