package community

import "time"

type Community struct {
	Id            string    `db:"id"`
	Name          string    `db:"name"`
	ImageUrl      string    `db:"imageUrl"`
	ImageId       string    `db:"imageId"`
	BackgroundUrl string    `db:"backgroundUrl"`
	BackgroundId  string    `db:"backgroundId"`
	Description   string    `db:"description"`
	Owner         string    `db:"owner"`
	CreatedAt     time.Time `db:"createdAt"`
	UpdatedAt     time.Time `db:"updatedAt"`
}
