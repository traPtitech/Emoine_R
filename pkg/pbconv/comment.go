package pbconv

import (
	"github.com/samber/lo"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBComment(c dbmodel.Comment) *emoine_rv1.Comment {
	return &emoine_rv1.Comment{
		Id:          c.ID.String(),
		EventId:     c.EventID.String(),
		Username:    lo.Ternary(c.IsAnonymous, "", c.UserID),
		Text:        c.Text,
		IsAnonymous: c.IsAnonymous,
		Color:       lo.Ternary(c.Color.Valid, c.Color.String, ""),
		CreatedAt:   timestamppb.New(c.CreatedAt),
	}
}
