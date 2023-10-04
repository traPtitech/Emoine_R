package handler

import (
	"context"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/pkg/pbconv"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
	"github.com/traPtitech/Emoine_R/repository"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
)

type GeneralAPIHandler struct {
	r      *repository.Repository
	logger *slog.Logger
}

func NewGeneralAPIHandler(r *repository.Repository, logger *slog.Logger) emoine_rv1connect.GeneralAPIServiceHandler {
	return &GeneralAPIHandler{
		r:      r,
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

	m, cnt, err := h.r.SelectEvents(ctx, int(*req.Msg.Limit), int(*req.Msg.Offset))
	if err != nil {
		h.logger.Error("SelectEvents", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("イベントの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetEventsResponse{
		Total: int32(cnt),
		Events: lo.Map(m, func(v dbmodel.Event, _ int) *emoine_rv1.Event {
			return pbconv.FromDBEvent(v)
		}),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetEvent(ctx context.Context, req *connect.Request[emoine_rv1.GetEventRequest]) (*connect.Response[emoine_rv1.GetEventResponse], error) {
	eid, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		h.logger.Error("Parse", "err", err)

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("eventIdのパースに失敗しました"))
	}

	e, err := h.r.SelectEvent(ctx, eid)
	if err != nil {
		h.logger.Error("SelectEvent", "err", err)

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

func (h *GeneralAPIHandler) GetEventReactions(ctx context.Context, req *connect.Request[emoine_rv1.GetEventReactionsRequest]) (*connect.Response[emoine_rv1.GetEventReactionsResponse], error) {
	eid, err := uuid.Parse(req.Msg.EventId)
	if err != nil {
		h.logger.Error("failed to parse event id", "err", err)

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("eventIDのパースに失敗しました"))
	}

	reactions, err := h.r.SelectEventReactions(ctx, eid)
	if err != nil {
		h.logger.Error("failed to select event reactions", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("リアクションの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetEventReactionsResponse{
		Reactions: lo.Map(reactions, func(r dbmodel.Reaction, _ int) *emoine_rv1.Reaction {
			return pbconv.FromDBReaction(r)
		}),
	})

	return res, nil
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
