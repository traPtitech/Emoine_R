package pbconv

import (
	"time"

	"github.com/traPtitech/Emoine_R/model/dbschema"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPBMeeting(m dbschema.Meeting) *emoine_rv1.Meeting {
	return &emoine_rv1.Meeting{
		Id:          m.ID.String(),
		VideoId:     m.VideoID,
		Title:       m.Title,
		Thumbnail:   m.Thumbnail,
		Description: mustValue[string](m.Description),
		StartedAt:   timestamppb.New(m.StartedAt),
		EndedAt:     timestamppb.New(mustValue[time.Time](m.EndedAt)),
	}
}
