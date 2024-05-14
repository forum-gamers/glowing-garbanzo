package member

import "time"

type Member struct {
	Id          string    `db:"id"`
	CommunityId string    `db:"communityId"`
	UserId      string    `db:"userId"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}
