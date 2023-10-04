package pbconv

import (
	"github.com/samber/lo"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBToken(t dbmodel.Token) *emoine_rv1.Token {
	return &emoine_rv1.Token{
		Id:          t.ID.String(),
		Raw:         t.Value,
		Username:    t.UserID,
		EventId:     t.EventID.String(),
		CreatorId:   t.CreatorID,
		Description: lo.Ternary(t.Description.Valid, t.Description.String, ""),
		CreatedAt:   timestamppb.New(t.CreatedAt),
		ExpireAt:    lo.Ternary(t.ExpireAt.Valid, timestamppb.New(t.ExpireAt.Time), nil),
	}
}
