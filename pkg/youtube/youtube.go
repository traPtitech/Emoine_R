package youtube

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var apiKey = os.Getenv("YOUTUBE_API_KEY")

func newYoutubeService(ctx context.Context) (*youtube.Service, error) {
	return youtube.NewService(ctx, option.WithAPIKey(apiKey))
}

func GetVideo(ctx context.Context, videoID string) (*youtube.Video, error) {
	service, err := newYoutubeService(ctx)
	if err != nil {
		return nil, err
	}

	call := service.Videos.List([]string{"snippet", "liveStreamingDetails"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("No video found for %s", videoID)
	}

	return response.Items[0], nil
}
