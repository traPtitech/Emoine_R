package pbconv

import (
	"github.com/samber/lo"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBEvent(e dbmodel.Event) *emoine_rv1.Event {
	return &emoine_rv1.Event{
		Id:          e.ID.String(),
		VideoId:     e.VideoID,
		Title:       e.Title,
		Thumbnail:   e.Thumbnail,
		Description: lo.Ternary(e.Description.Valid, e.Description.String, ""),
		StartedAt:   timestamppb.New(e.StartedAt),
		EndedAt:     lo.Ternary(e.EndedAt.Valid, timestamppb.New(e.EndedAt.Time), nil),
	}
}
