package pbconv

import (
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
	"github.com/traPtitech/Emoine_R/repository/dbmodel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBReaction(r dbmodel.Reaction) *emoine_rv1.Reaction {
	return &emoine_rv1.Reaction{
		Id:        r.ID.String(),
		EventId:   r.EventID.String(),
		Username:  r.UserID,
		StampId:   r.StampID.String(),
		CreatedAt: timestamppb.New(r.CreatedAt),
	}
}
