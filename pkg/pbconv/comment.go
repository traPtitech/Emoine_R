package pbconv

import (
	"github.com/traPtitech/Emoine_R/model/dbschema"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBComment(c dbschema.Comment) *emoine_rv1.Comment {
	return &emoine_rv1.Comment{
		Id:          c.ID.String(),
		MeetingId:   c.MeetingID.String(),
		Username:    c.UserID,
		Text:        c.Text,
		IsAnonymous: c.IsAnonymous,
		Color:       c.Color.String,
		CreatedAt:   timestamppb.New(c.CreatedAt),
	}
}
