package handler

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"google.golang.org/api/youtube/v3"
)

func mustValue[T any](v driver.Valuer) (value T) {
	if v == nil {
		return
	}

	vv, err := v.Value()
	if err != nil {
		return
	}

	value, _ = vv.(T)

	return value
}

func getVideoStreamingDates(video *youtube.Video) (time.Time, sql.NullTime, error) {
	startedAt, err := time.Parse(time.RFC3339, video.LiveStreamingDetails.ActualStartTime)
	if err != nil {
		return time.Time{}, sql.NullTime{}, err
	}

	var endedAt sql.NullTime
	if video.LiveStreamingDetails.ActualEndTime != "" {
		endedAt.Time, err = time.Parse(time.RFC3339, video.LiveStreamingDetails.ActualEndTime)
		if err != nil {
			return time.Time{}, sql.NullTime{}, err
		}
		endedAt.Valid = true
	}

	return startedAt, endedAt, nil
}
