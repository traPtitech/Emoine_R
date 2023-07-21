package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine_R/handler/schema"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	"github.com/traPtitech/Emoine_R/pkg/youtube"
)

func CreateMeeting(c echo.Context) error {
	req := new(schema.CreateMeetingJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストのパースに失敗しました").SetInternal(err)
	}

	c.Logger().Info(req.VideoId, req.Description)
	video, err := youtube.GetVideo(c.Request().Context(), req.VideoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "動画の取得に失敗しました").SetInternal(err)
	}

	startedAt, endedAt, err := getVideoStreamingDates(video)
	if err != nil {
		msg := "動画の開始時刻または終了時刻の取得に失敗しました"
		if errors.Is(err, errIsNotLiveStreaming) {
			msg += ": ライブ配信のIDを指定してください"
		}

		return echo.NewHTTPError(http.StatusBadRequest, msg).SetInternal(err)
	}

	var description sql.NullString
	if len(req.Description) > 0 {
		description.String = req.Description
		description.Valid = true
	}

	m := dbschema.Meeting{
		ID:          uuid.New(),
		VideoID:     req.VideoId,
		Title:       video.Snippet.Title,
		Thumbnail:   video.Snippet.Thumbnails.High.Url,
		Description: description,
		StartedAt:   startedAt,
		EndedAt:     endedAt,
	}
	if err := m.Insert(c.Request().Context(), model.DB); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ミーティングの作成に失敗しました").SetInternal(err)
	}

	return c.JSON(http.StatusCreated, schema.Meeting{
		Id:          m.ID,
		VideoId:     m.VideoID,
		Title:       m.Title,
		Thumbnail:   m.Thumbnail,
		Description: mustValue[string](m.Description),
		StartedAt:   m.StartedAt,
		EndedAt:     mustValue[time.Time](m.EndedAt),
	})
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

func UpdateMeeting(c echo.Context) error {
	req := new(schema.UpdateMeetingJSONRequestBody)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "リクエストのパースに失敗しました").SetInternal(err)
	}

	mid, err := uuid.Parse(c.Param("meetingId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "meetingIdのパースに失敗しました").SetInternal(err)
	}

	m, err := dbschema.MeetingByID(c.Request().Context(), model.DB, mid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ミーティングの取得に失敗しました").SetInternal(err)
	}

	if len(req.Description) > 0 {
		m.Description.String = req.Description
		m.Description.Valid = true
	}

	if err := m.Update(c.Request().Context(), model.DB); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "ミーティングの更新に失敗しました").SetInternal(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteMeeting(c echo.Context) error {
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
