package models

import (
	"time"

	"github.com/google/uuid"
)

type UserReaction struct {
	UserId    uuid.UUID `db:"user_id"`
	PostId    uuid.UUID `db:"post_id"`
	Reaction  string    `db:"reaction"`
	CreatedAt time.Time `db:"created_at"`
}
