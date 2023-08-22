package handler

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/pkg/youtube"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *AdminAPIHandler) CreateMeeting(ctx context.Context, req *connect.Request[emoine_rv1.CreateMeetingRequest]) (*connect.Response[emoine_rv1.CreateMeetingResponse], error) {
	h.logger.Info("Video Info", slog.String("video_id", req.Msg.VideoId), slog.String("description", req.Msg.Description))
	video, err := youtube.GetVideo(ctx, req.Msg.VideoId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("動画の取得に失敗しました"))
	}

	startedAt, endedAt, err := getVideoStreamingDates(video)
	if err != nil {
		msg := "動画の開始時刻または終了時刻の取得に失敗しました"
		if errors.Is(err, errIsNotLiveStreaming) {
			msg += ": ライブ配信のIDを指定してください"
		}

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New(msg))
	}

	var description sql.NullString
	if req.Msg.Description != "" {
		description.String = req.Msg.Description
		description.Valid = true
	}

	m := dbschema.Meeting{
		ID:          uuid.New(),
		VideoID:     req.Msg.VideoId,
		Title:       video.Snippet.Title,
		Thumbnail:   video.Snippet.Thumbnails.High.Url,
		Description: description,
		StartedAt:   startedAt,
		EndedAt:     endedAt,
	}
	if err := m.Insert(ctx, model.DB); err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの作成に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.CreateMeetingResponse{
		Meeting: &emoine_rv1.Meeting{
			Id:          m.ID.String(),
			VideoId:     m.VideoID,
			Title:       m.Title,
			Thumbnail:   m.Thumbnail,
			Description: m.Description.String,
			StartedAt:   timestamppb.New(m.StartedAt),
			EndedAt:     lo.Ternary(m.EndedAt.Valid, timestamppb.New(m.EndedAt.Time), nil),
		},
	})

	return res, nil
}

func (h *AdminAPIHandler) UpdateMeeting(ctx context.Context, req *connect.Request[emoine_rv1.UpdateMeetingRequest]) (*connect.Response[emptypb.Empty], error) {
	mid, err := uuid.Parse(req.Msg.MeetingId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("meetingIdのパースに失敗しました"))
	}

	m, err := dbschema.MeetingByID(ctx, model.DB, mid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}

	if req.Msg.Description != nil {
		m.Description.String = *req.Msg.Description
		m.Description.Valid = true
	}

	if err := m.Update(ctx, model.DB); err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの更新に失敗しました"))
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (h *AdminAPIHandler) DeleteMeeting(ctx context.Context, req *connect.Request[emoine_rv1.DeleteMeetingRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

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
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}
	cnt, err := dbschema.MeetingCount(ctx, model.DB)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetMeetingsResponse{
		Total: int32(cnt),
		Meetings: lo.Map(m, func(v dbschema.Meeting, _ int) *emoine_rv1.Meeting {
			return &emoine_rv1.Meeting{
				Id:          v.ID.String(),
				VideoId:     v.VideoID,
				Title:       v.Title,
				Thumbnail:   v.Thumbnail,
				Description: v.Description.String,
				StartedAt:   timestamppb.New(v.StartedAt),
				EndedAt:     lo.Ternary(v.EndedAt.Valid, timestamppb.New(v.EndedAt.Time), nil),
			}
		}),
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetMeeting(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingRequest]) (*connect.Response[emoine_rv1.GetMeetingResponse], error) {
	mid, err := uuid.Parse(req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("meetingIdのパースに失敗しました"))
	}

	m, err := dbschema.MeetingByID(ctx, model.DB, mid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.GetMeetingResponse{
		Meeting: &emoine_rv1.Meeting{
			Id:          m.ID.String(),
			VideoId:     m.VideoID,
			Title:       m.Title,
			Thumbnail:   m.Thumbnail,
			Description: m.Description.String,
			StartedAt:   timestamppb.New(m.StartedAt),
			EndedAt:     lo.Ternary(m.EndedAt.Valid, timestamppb.New(m.EndedAt.Time), nil),
		},
	})

	return res, nil
}

func (h *GeneralAPIHandler) GetMeetingComments(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingCommentsRequest]) (*connect.Response[emoine_rv1.GetMeetingCommentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *GeneralAPIHandler) GetMeetingReactions(ctx context.Context, req *connect.Request[emoine_rv1.GetMeetingReactionsRequest]) (*connect.Response[emoine_rv1.GetMeetingReactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

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

func GetMeetingTokens(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
