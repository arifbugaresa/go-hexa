package aktivitas

import "time"

type Aktivitas struct {
	ID         int64
	UserInfoID int64
	Name       string
	CreatedBy  string
	CreatedAt  time.Time
	UpdatedBy  string
	UpdatedAt  time.Time
	Deleted    bool
}
