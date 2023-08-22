package handler

import (
	"database/sql"
	"time"

	"google.golang.org/api/youtube/v3"
)

func getVideoStreamingDates(video *youtube.Video) (time.Time, sql.NullTime, error) {
	videoLiveStreamingDetails := video.LiveStreamingDetails
	if videoLiveStreamingDetails == nil {
		return time.Time{}, sql.NullTime{}, errIsNotLiveStreaming
	}

	startedAt, err := time.Parse(time.RFC3339, videoLiveStreamingDetails.ActualStartTime)
	if err != nil {
		return time.Time{}, sql.NullTime{}, err
	}

	var endedAt sql.NullTime
	if videoLiveStreamingDetails.ActualEndTime != "" {
		endedAt.Time, err = time.Parse(time.RFC3339, video.LiveStreamingDetails.ActualEndTime)
		if err != nil {
			return time.Time{}, sql.NullTime{}, err
		}
		endedAt.Valid = true
	}

	return startedAt, endedAt, nil
}
