package handler

import (
	"context"
	"errors"
	"log/slog"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	"github.com/traPtitech/Emoine_R/pkg/pbconv"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
)

type GeneralAPIHandler struct {
	logger *slog.Logger
}

func NewGeneralAPIHandler(logger *slog.Logger) emoine_rv1connect.GeneralAPIServiceHandler {
	return &GeneralAPIHandler{
		logger: logger,
	}
}

func (h *GeneralAPIHandler) GetMeetings(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingsRequest]) (*connect.Response[emoine_rv1.GetMeetingsResponse], error) {
	if req.Msg.Limit == nil {
		limit := int32(10)
		req.Msg.Limit = &limit
	}
	if req.Msg.Offset == nil {
		offset := int32(0)
		req.Msg.Offset = &offset
	}
	m, err := dbschema.Meetings(ctx, model.DB, int(*req.Msg.Limit), int(*req.Msg.Offset))
	if err != nil {
		h.logger.Error("Meetings", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}
	cnt, err := dbschema.MeetingCount(ctx, model.DB)
	if err != nil {
		h.logger.Error("MeetingCount", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetMeetingsResponse{
		Total: int32(cnt),
		Meetings: lo.Map(m, func(v dbschema.Meeting, _ int) *emoine_rv1.Meeting {
			return pbconv.FromDBMeeting(v)
		}),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetMeeting(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingRequest]) (*connect.Response[emoine_rv1.GetMeetingResponse], error) {
	mid, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		h.logger.Error("Parse", "err", err)

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("meetingIdのパースに失敗しました"))
	}

	m, err := dbschema.MeetingByID(ctx, model.DB, mid)
	if err != nil {
		h.logger.Error("MeetingByID", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}
	if m == nil {
		h.logger.Error("MeetingByID", "err", "not found")

		return nil, connect.NewError(connect.CodeNotFound, errors.New("ミーティングが見つかりませんでした"))
	}

	res := connect.NewResponse(&emoine_rv1.GetMeetingResponse{
		Meeting: pbconv.FromDBMeeting(*m),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetMeetingComments(_ context.Context, _ *connect.Request[emoine_rv1.GetMeetingCommentsRequest]) (*connect.Response[emoine_rv1.GetMeetingCommentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) GetMeetingReactions(_ context.Context, _ *connect.Request[emoine_rv1.GetMeetingReactionsRequest]) (*connect.Response[emoine_rv1.GetMeetingReactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) ConnectToMeetingStream(
	_ context.Context,
	_ *connect.Request[emoine_rv1.ConnectToMeetingStreamRequest],
	_ *connect.ServerStream[emoine_rv1.ConnectToMeetingStreamResponse],
) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendComment(_ context.Context, _ *connect.Request[emoine_rv1.SendCommentRequest]) (*connect.Response[emoine_rv1.SendCommentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendReaction(_ context.Context, _ *connect.Request[emoine_rv1.SendReactionRequest]) (*connect.Response[emoine_rv1.SendReactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}
