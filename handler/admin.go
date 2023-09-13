package handler

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	"github.com/traPtitech/Emoine_R/pkg/pbconv"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1/emoine_rv1connect"
	"github.com/traPtitech/Emoine_R/pkg/youtube"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AdminAPIHandler struct {
	logger *slog.Logger
}

func NewAdminAPIHandler(logger *slog.Logger) emoine_rv1connect.AdminAPIServiceHandler {
	return &AdminAPIHandler{
		logger: logger,
	}
}

func (h *AdminAPIHandler) CreateMeeting(ctx context.Context, req *connect.Request[emoine_rv1.CreateMeetingRequest]) (*connect.Response[emoine_rv1.CreateMeetingResponse], error) {
	video, err := youtube.GetVideo(ctx, req.Msg.VideoId)
	if err != nil {
		h.logger.Error("GetVideo", "err", err)

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("動画の取得に失敗しました"))
	}

	startedAt, endedAt, err := getVideoStreamingDates(video)
	if err != nil {
		msg := "動画の開始時刻または終了時刻の取得に失敗しました"
		if errors.Is(err, errIsNotLiveStreaming) {
			msg += ": ライブ配信のIDを指定してください"
		}

		h.logger.Error("GetVideoStreamingDates", "err", err)

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
		h.logger.Error("Insert", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの作成に失敗しました"))
	}

	res := connect.NewResponse(&emoine_rv1.CreateMeetingResponse{
		Meeting: pbconv.FromDBMeeting(m),
	})

	return res, nil
}

func (h *AdminAPIHandler) UpdateMeeting(ctx context.Context, req *connect.Request[emoine_rv1.UpdateMeetingRequest]) (*connect.Response[emptypb.Empty], error) {
	mid, err := uuid.Parse(req.Msg.MeetingId)
	if err != nil {
		h.logger.Error("Parse", "err", err)

		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("meetingIdのパースに失敗しました"))
	}

	m, err := dbschema.MeetingByID(ctx, model.DB, mid)
	if err != nil {
		h.logger.Error("MeetingByID", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの取得に失敗しました"))
	}

	if req.Msg.Description != nil {
		m.Description.String = *req.Msg.Description
		m.Description.Valid = true
	}

	if err := m.Update(ctx, model.DB); err != nil {
		h.logger.Error("Update", "err", err)

		return nil, connect.NewError(connect.CodeInternal, errors.New("ミーティングの更新に失敗しました"))
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (h *AdminAPIHandler) DeleteMeeting(_ context.Context, _ *connect.Request[emoine_rv1.DeleteMeetingRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *AdminAPIHandler) GetTokens(_ context.Context, _ *connect.Request[emoine_rv1.GetTokensRequest]) (*connect.Response[emoine_rv1.GetTokensResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *AdminAPIHandler) GenerateToken(_ context.Context, _ *connect.Request[emoine_rv1.GenerateTokenRequest]) (*connect.Response[emoine_rv1.GenerateTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func (h *AdminAPIHandler) RevokeToken(_ context.Context, _ *connect.Request[emoine_rv1.RevokeTokenRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}
