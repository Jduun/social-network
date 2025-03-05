package models

import (
	"github.com/google/uuid"
)

type Follower struct {
	FollowerId uuid.UUID `db:"follower_id"`
	Following uuid.UUID `db:"following_id"`
}
