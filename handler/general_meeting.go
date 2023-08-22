package handler

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	"github.com/traPtitech/Emoine_R/pkg/pbconv"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
)

func (h *GeneralAPIHandler) GetMeetings(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingsRequest]) (*connect.Response[emoine_rv1.GetMeetingsResponse], error) {
	if req.Msg.Limit == nil {
		limit := int32(0)
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
			return pbconv.ToPBMeeting(v)
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

	res := connect.NewResponse(&emoine_rv1.GetMeetingResponse{
		Meeting: pbconv.ToPBMeeting(*m),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetMeetingComments(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingCommentsRequest]) (*connect.Response[emoine_rv1.GetMeetingCommentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) GetMeetingReactions(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingReactionsRequest]) (*connect.Response[emoine_rv1.GetMeetingReactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}
