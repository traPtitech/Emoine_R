package handler

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/traPtitech/Emoine_R/handler/schema"
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

func GetMeetings(c echo.Context) error {
	req := new(schema.GetMeetingsParams)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストのパースに失敗しました").SetInternal(err)
	}
	if req.Limit == nil {
		limit := 0
		req.Limit = &limit
	}
	if req.Offset == nil {
		offset := 0
		req.Offset = &offset
	}
	m, err := dbschema.Meetings(c.Request().Context(), model.DB, *req.Limit, *req.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ミーティングの取得に失敗しました").SetInternal(err)
	}
	cnt, err := dbschema.MeetingCount(c.Request().Context(), model.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ミーティングの取得に失敗しました").SetInternal(err)
	}

	return c.JSON(http.StatusOK, schema.MeetingsWithTotal{
		Total: cnt,
		Meetings: func() []schema.Meeting {
			ms := make([]schema.Meeting, len(m))
			for i, v := range m {
				ms[i] = schema.Meeting{
					Id:          v.ID,
					VideoId:     v.VideoID,
					Title:       v.Title,
					Thumbnail:   v.Thumbnail,
					Description: mustValue[string](v.Description),
					StartedAt:   v.StartedAt,
					EndedAt:     mustValue[time.Time](v.EndedAt),
				}
			}

			return ms
		}(),
	})
}

func GetMeeting(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetMeetingComments(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetMeetingReactions(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetMeetingTokens(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
