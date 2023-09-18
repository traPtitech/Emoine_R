package handler

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
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

func (h *GeneralAPIHandler) GetEvents(ctx context.Context, req *connect.Request[emoine_rv1.GetEventsRequest]) (*connect.Response[emoine_rv1.GetEventsResponse], error) {
	if req.Msg.Limit == nil {
		limit := int32(10)
		req.Msg.Limit = &limit
	}
	if req.Msg.Offset == nil {
		offset := int32(0)
		req.Msg.Offset = &offset
	}
	e, err := dbschema.Events(ctx, model.DB, int(*req.Msg.Limit), int(*req.Msg.Offset))
	if err != nil {
		h.logger.Error("Events", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("イベントの取得に失敗しました"))
	}
	cnt, err := dbschema.EventCount(ctx, model.DB)
	if err != nil {
		h.logger.Error("EventCount", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("イベントの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetEventsResponse{
		Total: int32(cnt),
		Events: lo.Map(e, func(v dbschema.Event, _ int) *emoine_rv1.Event {
			return pbconv.FromDBEvent(v)
		}),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetEvent(ctx context.Context, req *connect.Request[emoine_rv1.GetEventRequest]) (*connect.Response[emoine_rv1.GetEventResponse], error) {
	eid, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, withErrInfo(connect.CodeInvalidArgument, err, "eventIdのパースに失敗しました")
	}

	e, err := dbschema.EventByID(ctx, model.DB, eid)
	if err != nil {
		h.logger.Error("EventByID", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("イベントの取得に失敗しました"))
	}
	if e == nil {
		h.logger.Error("EventByID", "err", "not found")

		return nil, connect.NewError(connect.CodeNotFound, errors.New("イベントが見つかりませんでした"))
	}

	res := connect.NewResponse(&emoine_rv1.GetEventResponse{
		Event: pbconv.FromDBEvent(*e),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetEventComments(_ context.Context, _ *connect.Request[emoine_rv1.GetEventCommentsRequest]) (*connect.Response[emoine_rv1.GetEventCommentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) GetEventReactions(_ context.Context, _ *connect.Request[emoine_rv1.GetEventReactionsRequest]) (*connect.Response[emoine_rv1.GetEventReactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) ConnectToEventStream(
	_ context.Context,
	_ *connect.Request[emoine_rv1.ConnectToEventStreamRequest],
	_ *connect.ServerStream[emoine_rv1.ConnectToEventStreamResponse],
) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendComment(_ context.Context, _ *connect.Request[emoine_rv1.SendCommentRequest]) (*connect.Response[emoine_rv1.SendCommentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) SendReaction(_ context.Context, _ *connect.Request[emoine_rv1.SendReactionRequest]) (*connect.Response[emoine_rv1.SendReactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}
